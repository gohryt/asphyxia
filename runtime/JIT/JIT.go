package JIT

import (
	"syscall"
	"unsafe"
)

type (
	builder struct {
		operands []uint16
	}
)

func Builder() *builder {
	return new(builder)
}

func Assemble[T_function any](builder *builder, function *T_function) error {
	operands := builder.operands
	l := len(operands) * 2

	executable, err := syscall.Mmap(-1, 0, l, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_PRIVATE|syscall.MAP_ANONYMOUS)
	if err != nil {
		return err
	}

	j := 0

	for _, operand := range operands {
		executable[j] = byte(operand >> 8)
		executable[j+1] = byte(operand)
		j = j + 2
	}

	functionPointer := (uintptr)(unsafe.Pointer(&executable))
	*function = *(*T_function)(unsafe.Pointer(&functionPointer))

	return nil
}
