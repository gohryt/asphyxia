package runtime

import (
	"runtime"
	"unsafe"
)

type (
	Slice struct {
		Data     unsafe.Pointer
		Length   int
		Capacity int
	}

	String struct {
		Data   unsafe.Pointer
		Length int
	}

	InterfaceEmpty struct {
		Type unsafe.Pointer
		Data unsafe.Pointer
	}
)

var (
	GOOS   = runtime.GOOS
	GOARCH = runtime.GOARCH
)

func As[T_as, T_from *any](from *T_from) *T_as {
	return (*T_as)(unsafe.Pointer(from))
}
