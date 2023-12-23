package context

import (
	"time"
	"unsafe"
)

type (
	Context struct {
		Context unsafe.Pointer
		Done    func(context unsafe.Pointer) chan struct{}
		Cancel  func(context unsafe.Pointer)
	}

	WithCancelImplementation struct {
		done chan struct{}
	}

	WithTimeoutImplementation struct {
		done  chan struct{}
		timer *time.Timer
	}
)

func BackgroundDone(unsafe.Pointer) chan struct{} {
	return nil
}

func BackgroundCancel(unsafe.Pointer) {
}

func Background() *Context {
	return &Context{
		Context: unsafe.Pointer(nil),
		Done:    BackgroundDone,
		Cancel:  BackgroundCancel,
	}
}

func WithCancelDone(context unsafe.Pointer) chan struct{} {
	return (*WithCancelImplementation)(context).done
}

func WithCancelCancel(context unsafe.Pointer) {
	close((*WithCancelImplementation)(context).done)
}

func WithCancel() *Context {
	context := &WithCancelImplementation{
		done: make(chan struct{}),
	}

	return &Context{
		Context: unsafe.Pointer(context),
		Done:    WithCancelDone,
		Cancel:  WithCancelCancel,
	}
}

func WithTimeoutDone(context unsafe.Pointer) chan struct{} {
	return (*WithTimeoutImplementation)(context).done
}

func WithTimeoutCancel(context unsafe.Pointer) {
	close((*WithTimeoutImplementation)(context).done)
	(*WithTimeoutImplementation)(context).timer.Stop()
}

func WithTimeout(timeout time.Duration) *Context {
	done := make(chan struct{})

	timer := time.AfterFunc(timeout, func() {
		close(done)
	})

	context := &WithTimeoutImplementation{
		done:  done,
		timer: timer,
	}

	return &Context{
		Context: unsafe.Pointer(context),
		Done:    WithTimeoutDone,
		Cancel:  WithTimeoutCancel,
	}
}
