package binary_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/gohryt/asphyxia/encoding/binary"
)

var (
	BenchmarkCopyFrom     = make([]uint16, math.MaxUint16)
	BenchmarkCopyExpected = make([]byte, (math.MaxUint16 * 2))
)

func init() {
	i := 0
	j := 0

	for i < math.MaxUint16 {
		element := uint16(i)

		BenchmarkCopyFrom[i] = element

		BenchmarkCopyExpected[j] = byte(element >> 8)
		BenchmarkCopyExpected[j+1] = byte(element)

		i += 1
		j += 2
	}
}

func TestCopyNative(t *testing.T) {
	to := make([]byte, len(BenchmarkCopyFrom)*2)

	count := binary.BigEndianCopyUint16(to, BenchmarkCopyFrom...)

	if count != len(to) {
		t.Fatal("expected length is", len(to), "result is", count)
	}

	if !bytes.Equal(to, BenchmarkCopyExpected) {
		t.Fatal("expected result is", BenchmarkCopyExpected, "result is", to)
	}
}

func BenchmarkCopyNative(b *testing.B) {
	to := make([]byte, len(BenchmarkCopyFrom)*2)

	binary.BigEndianCopyUint16(to, BenchmarkCopyFrom...)
}

func BenchmarkCopy(b *testing.B) {
	to := make([]byte, len(BenchmarkCopyFrom)*2)

	BigEndianCopyUint16(to, BenchmarkCopyFrom...)
}

func BigEndianCopyUint16(to []byte, from ...uint16) int {
	i := 0
	j := 0

	count := len(to) / 2

	if len(from) < count {
		count = len(from)
	}

	for i < count {
		element := from[i]

		to[j] = byte(element >> 8)
		to[j+1] = byte(element)

		i += 1
		j += 2
	}

	return j
}
