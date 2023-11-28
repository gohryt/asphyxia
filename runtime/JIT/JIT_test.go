package JIT_test

import (
	"testing"

	"github.com/gohryt/asphyxia/runtime/JIT"
)

var test func()

func TestAssemble(t *testing.T) {
	builder := JIT.Builder()

	JIT.Assemble(builder, &test)

	test()
}
