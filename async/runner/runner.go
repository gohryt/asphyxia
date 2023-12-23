package runner

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
	"unsafe"
)

type (
	Action uint64

	Group struct {
		shutdown context.Context
		cancel   context.CancelFunc

		onStart []todo
		onClose []todo

		wait uint64
	}

	Task[T_resource any] func(ctx context.Context, group *Group, resource *T_resource)

	task func(ctx context.Context, group *Group, to unsafe.Pointer)

	todo struct {
		to unsafe.Pointer
		do task
	}
)

const (
	ActionNone Action = iota + 0
	ActionStart
	ActionClose
)

func New(parent context.Context, signals ...os.Signal) (group *Group, shutdown context.Context) {
	group = new(Group)
	group.shutdown, group.cancel = signal.NotifyContext(context.Background(), signals...)

	return group, group.shutdown
}

func On[T_resource any](action Action, group *Group, resource *T_resource, function Task[T_resource]) {
	todo := todo{
		to: unsafe.Pointer(resource),
		do: *(*task)(unsafe.Pointer(&function)),
	}

	switch action {
	case ActionStart:
		group.onStart = append(group.onStart, todo)
	case ActionClose:
		group.onClose = append(group.onClose, todo)
	}
}

func Send(group *Group, action Action) {
	if action == ActionClose {
		group.cancel()
	}
}

func wrap(ctx context.Context, group *Group, todo todo) {
	atomic.AddUint64(&group.wait, 1)

	todo.do(ctx, group, todo.to)

	if atomic.AddUint64(&group.wait, ^uint64(0)) == 0 {
		group.cancel()
	}
}

func Wait(group *Group, timeout time.Duration) {
	for _, todo := range group.onStart {
		go wrap(group.shutdown, group, todo)
	}

	<-group.shutdown.Done()

	background := context.Background()

	if timeout > 0 {
		group.shutdown, group.cancel = context.WithTimeout(background, timeout)
	} else {
		group.shutdown, group.cancel = context.WithCancel(background)
	}

	for _, todo := range group.onClose {
		go wrap(group.shutdown, group, todo)
	}

	<-group.shutdown.Done()
}
