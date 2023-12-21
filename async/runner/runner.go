package runner

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
)

type (
	Action uint64

	Group struct {
		shutdown context.Context
		cancel   context.CancelFunc
		actions  chan task

		locker  sync.Mutex
		onStart []task
		onClose []task
		on      sync.Map

		registered uint32
		async      bool
	}

	Task[T_resource any] func(ctx context.Context, group *Group, resource *T_resource) error

	wrapper func(ctx context.Context, group *Group, resource, function any) error

	task struct {
		resource any
		function any
		wrapper  wrapper
	}
)

const (
	ActionNone Action = iota + 0
	ActionStart
	ActionClose
)

func Prepare(parent context.Context, signals ...os.Signal) (group *Group, shutdown context.Context) {
	group = &Group{
		registered: uint32(ActionClose),
	}

	if len(signals) > 0 {
		group.shutdown, group.cancel = signal.NotifyContext(parent, signals...)
	} else {
		group.shutdown, group.cancel = context.WithCancel(parent)
	}

	return group, group.shutdown
}

func On[T_resource any](action Action, group *Group, resource *T_resource, function Task[T_resource]) {
	task := task{
		resource: resource,
		function: function,
		wrapper:  wrap[T_resource],
	}

	if group.async {
		group.locker.Lock()
		defer group.locker.Unlock()
	}

	switch action {
	case ActionStart:
		group.onStart = append(group.onStart, task)
	case ActionClose:
		group.onClose = append(group.onClose, task)
	}
}

func wrap[T_resource any](ctx context.Context, group *Group, resource, function any) error {
	return function.(Task[T_resource])(ctx, group, resource.(*T_resource))
}

func (group *Group) Register() uint32 {
	return atomic.AddUint32(&group.registered, 1)
}

func (group *Group) Wait() {
	shutdown := group.shutdown

	for _, task := range group.onStart {
		go task.wrapper(shutdown, group, task.resource, task.function)
	}

	select {
	case <-shutdown.Done():
	}

	ctx := context.Background()

	for _, task := range group.onClose {
		go task.wrapper(ctx, group, task.resource, task.function)
	}
}
