package JIT_test

import (
	"testing"

	"github.com/gohryt/asphyxia/runtime/JIT"
)

var test func()

func TestAssemble(t *testing.T) {
	builder := JIT.Builder()

	builder.MOV(JIT.AX, 1)
	builder.RET()

	JIT.Assemble(builder, &test)

	test()
}
