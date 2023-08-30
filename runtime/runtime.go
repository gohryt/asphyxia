package runtime

import (
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

	Any struct {
		Type unsafe.Pointer
		Data unsafe.Pointer
	}
)

func As[AS, FROM any](from *FROM) *AS {
	return (*AS)(unsafe.Pointer(from))
}
