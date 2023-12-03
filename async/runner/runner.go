package runner

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
)

type (
	Action uint64

	Group struct {
		shutdown context.Context
		cancel   context.CancelFunc
		actions  chan task
		storage  map[Action][]task

		registered uint64
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
	ActionNone Action = 0 + iota
	ActionStart
	ActionClose
)

func Prepare(parent context.Context, signals ...os.Signal) (group *Group, shutdown context.Context) {
	group = &Group{
		storage:    make(map[Action][]task),
		registered: uint64(ActionClose),
	}

	if len(signals) > 0 {
		group.shutdown, group.cancel = signal.NotifyContext(parent, signals...)
	} else {
		group.shutdown, group.cancel = context.WithCancel(parent)
	}

	return group, group.shutdown
}

func On[T_resource any](action Action, group *Group, resource *T_resource, function Task[T_resource]) {
	group.storage[action] = append(group.storage[action], task{
		resource: resource,
		function: function,
		wrapper:  wrap[T_resource],
	})
}

func wrap[T_resource any](ctx context.Context, group *Group, resource, function any) error {
	return function.(Task[T_resource])(ctx, group, resource.(*T_resource))
}

func (group *Group) Register() uint64 {
	return atomic.AddUint64(&group.registered, 1)
}

func (group *Group) Wait() {
	shutdown := group.shutdown

	start := group.storage[ActionStart]

	for _, task := range start {
		go task.wrapper(group.shutdown, group, task.resource, task.function)
	}

	select {
	case task := <-group.actions:
		go task.wrapper(group.shutdown, group, task.resource, task.function)
	case <-shutdown.Done():
	}
}
