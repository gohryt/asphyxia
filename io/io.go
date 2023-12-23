package io

import (
	"io"
	"unsafe"
)

type (
	Reader struct {
		From unsafe.Pointer
		Read func(from unsafe.Pointer, to []byte) (n int, err error)
	}

	Writer struct {
		To    unsafe.Pointer
		Write func(to unsafe.Pointer, from []byte) (n int, err error)
	}
)

var EOF = io.EOF
