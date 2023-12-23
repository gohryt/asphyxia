package runner

import (
	"os"
	"os/signal"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/gohryt/asphyxia/async/context"
)

type (
	Action uint64

	Group struct {
		shutdown *context.Context

		onStart []todo
		onClose []todo

		wait uint64
	}

	Task[T_resource any] func(ctx *context.Context, group *Group, resource *T_resource)

	task func(ctx *context.Context, group *Group, to unsafe.Pointer)

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

func New(parent *context.Context, signals ...os.Signal) (group *Group, shutdown *context.Context) {
	group = &Group{
		shutdown: context.WithCancel(),
	}

	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, signals...)

	go func() {
		select {
		case <-shutdown.Done(shutdown.Context):
		case <-notifier:
			shutdown.Cancel(shutdown.Context)
		}
	}()

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
		group.shutdown.Cancel(group.shutdown.Context)
	}
}

func wrap(ctx *context.Context, group *Group, todo todo) {
	atomic.AddUint64(&group.wait, 1)

	todo.do(ctx, group, todo.to)

	if atomic.AddUint64(&group.wait, ^uint64(0)) == 0 {
		group.shutdown.Cancel(group.shutdown.Context)
	}
}

func Wait(group *Group, timeout time.Duration) {
	for _, todo := range group.onStart {
		go wrap(group.shutdown, group, todo)
	}

	<-group.shutdown.Done(group.shutdown.Context)

	if timeout > 0 {
		group.shutdown = context.WithTimeout(timeout)
	} else {
		group.shutdown = context.WithCancel()
	}

	for _, todo := range group.onClose {
		go wrap(group.shutdown, group, todo)
	}

	<-group.shutdown.Done(group.shutdown.Context)
}
