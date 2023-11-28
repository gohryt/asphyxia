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
	return &builder{
		operands: []uint16{
			0x48c7, 0xc001, 0x0, // mov %rax,$0x1
			0x48, 0xc7c7, 0x100, 0x0, // mov %rdi,$0x1
			0x48c7, 0xc20c, 0x0, // mov 0x13, %rdx
			0x48, 0x8d35, 0x400, 0x0, // lea 0x4(%rip), %rsi
			0xf05,                  // syscall
			0xc3cc,                 // ret
			0x4865, 0x6c6c, 0x6f20, // Hello_(whitespace)
			0x576f, 0x726c, 0x6421, 0xa, // World!
		},
	}
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
