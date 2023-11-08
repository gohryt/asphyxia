package runtime_test

import (
	"testing"

	"github.com/gohryt/asphyxia-core/runtime"
)

type (
	TestAsA struct {
		x, y int
	}

	TestAsB [2]int
)

func TestAs(t *testing.T) {
	a := &TestAsA{x: 10, y: 10}
	b := runtime.As[TestAsB](a)

	if b[0] != 10 || b[1] != 10 {
		t.Fail()
	}
}
