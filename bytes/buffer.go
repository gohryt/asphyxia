package bytes

import (
	"io"
	"unicode/utf8"

	"github.com/gohryt/asphyxia/memory"
)

type (
	Buffer []byte
)

func (buffer *Buffer) Reset() {
	*buffer = (*buffer)[:0]
}

func (buffer *Buffer) String() string {
	return string(*buffer)
}

func (buffer *Buffer) Clone() *Buffer {
	clone := make(Buffer, len(*buffer))
	copy(clone, *buffer)
	return &clone
}

func (buffer *Buffer) Set(value []byte) {
	*buffer = append((*buffer)[:0], value...)
}

func (buffer *Buffer) SetString(value string) {
	*buffer = append((*buffer)[:0], value...)
}

func (buffer *Buffer) Write(value []byte) (n int, err error) {
	*buffer = append(*buffer, value...)
	return len(value), nil
}

func (buffer *Buffer) WriteString(value string) (n int, err error) {
	*buffer = append(*buffer, value...)
	return len(value), nil
}

func (buffer *Buffer) WriteByte(value byte) (err error) {
	*buffer = append(*buffer, value)
	return
}

func (buffer *Buffer) WriteRune(value rune) (n int, err error) {
	slice := *buffer
	l := len(slice)

	size := l + utf8.UTFMax

	if size > cap(slice) {
		reallocation := make(Buffer, size)
		copy(reallocation, slice)

		slice = reallocation
	}

	n = utf8.EncodeRune(slice[l:size], value)

	*buffer = slice[:(l + n)]
	return
}

func (buffer *Buffer) ReadFrom(from io.Reader) (n int64, err error) {
	slice := *buffer
	l := len(slice)
	r := 0

reallocation:
	size := cap(slice)

	if size == 0 {
		size = 64
	} else if size < (8 * memory.Kilobyte) {
		size *= 4
	} else {
		size *= 2
	}

	reallocation := make(Buffer, size)
	copy(reallocation, slice)

	slice = reallocation

read:
	r, err = from.Read(slice[l:size])

	n += int64(r)
	l += r

	if err == nil {
		if l < size {
			goto read
		}

		goto reallocation
	} else if err == io.EOF {
		err = nil
	}

	*buffer = slice[:l]
	return
}
