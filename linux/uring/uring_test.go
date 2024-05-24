package uring_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/gohryt/asphyxia/linux/uring"
)

func TestXxx(t *testing.T) {
	fmt.Println(unsafe.Sizeof(uring.IOURingSQE{}))
}
