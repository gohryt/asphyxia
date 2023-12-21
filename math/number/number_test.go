package number_test

import (
	"testing"

	"github.com/gohryt/asphyxia/math/number"
)

func TestAdd(t *testing.T) {
	result := number.Add("1", "1")

	if result != "2" {
		t.Fatalf("result should be %s not %s", "2", result)
	}
}
