package bytes_test

import (
	"testing"

	"github.com/gohryt/asphyxia-core/bytes"
)

var ExampleString = "example"

func TestReset(t *testing.T) {
	b := bytes.Buffer(ExampleString)
	b.Reset()

	if len(b) > 0 {
		t.Fatal(len(b))
	}
}

func TestAsString(t *testing.T) {
	b := bytes.Buffer(ExampleString)
	s := b.AsString()

	if s != ExampleString {
		t.Fatal(s)
	}
}