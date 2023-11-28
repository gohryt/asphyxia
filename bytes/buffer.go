package bytes

import (
	"io"
	"unicode/utf8"

	"github.com/gohryt/asphyxia/memory"
)

type (
	Buffer []byte
)

func (buffer *Buffer) Clone() *Buffer {
	clone := make(Buffer, len(*buffer))
	copy(clone, *buffer)
	return &clone
}

func (buffer *Buffer) Set(source []byte) {
	*buffer = append((*buffer)[:0], source...)
}

func (buffer *Buffer) SetString(source string) {
	*buffer = append((*buffer)[:0], source...)
}

func (buffer *Buffer) Write(source []byte) (n int, err error) {
	*buffer = append(*buffer, source...)
	return len(source), nil
}

func (buffer *Buffer) WriteByte(source byte) (err error) {
	*buffer = append(*buffer, source)
	return
}

func (buffer *Buffer) WriteRune(source rune) (n int, err error) {
	slice := *buffer
	l := len(slice)

	size := l + utf8.UTFMax

	if size > cap(slice) {
		reallocation := make(Buffer, size)
		copy(reallocation, slice)

		slice = reallocation
	}

	n = utf8.EncodeRune(slice[l:size], source)

	*buffer = slice[:(l + n)]
	return
}

func (buffer *Buffer) WriteString(source string) (n int, err error) {
	*buffer = append(*buffer, source...)
	return len(source), nil
}

func (buffer *Buffer) ReadFrom(source io.Reader) (n int64, err error) {
	slice := *buffer
	l := len(slice)
	r := 0

reallocation:
	size := cap(slice)

	if size == 0 {
		size = 64
	} else if size < memory.Kilobyte {
		size *= 4
	} else {
		size *= 2
	}

	reallocation := make(Buffer, size)
	copy(reallocation, slice)

	slice = reallocation

read:
	r, err = source.Read(slice[l:size])

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
