package JIT_test

import (
	"testing"

	"github.com/gohryt/asphyxia-core/runtime/JIT"
)

var test func()

func TestAssemble(t *testing.T) {
	builder := JIT.Builder("x")

	JIT.Assemble[func()](builder, &test)

	test()
}
