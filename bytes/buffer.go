package bytes

import (
	"unicode/utf8"

	"github.com/gohryt/asphyxia/io"
	"github.com/gohryt/asphyxia/memory"
)

type (
	Buffer []byte
)

func Reset(buffer *Buffer) {
	*buffer = (*buffer)[:0]
}

func String(buffer *Buffer) string {
	return string(*buffer)
}

func Clone(buffer *Buffer) *Buffer {
	clone := make(Buffer, len(*buffer))
	copy(clone, *buffer)
	return &clone
}

func Set(buffer *Buffer, value []byte) {
	*buffer = append((*buffer)[:0], value...)
}

func SetString(buffer *Buffer, value string) {
	*buffer = append((*buffer)[:0], value...)
}

func Write(buffer *Buffer, value []byte) (n int, err error) {
	*buffer = append(*buffer, value...)
	return len(value), nil
}

func WriteString(buffer *Buffer, value string) (n int, err error) {
	*buffer = append(*buffer, value...)
	return len(value), nil
}

func WriteByte(buffer *Buffer, value byte) (err error) {
	*buffer = append(*buffer, value)
	return
}

func WriteRune(buffer *Buffer, value rune) (n int, err error) {
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

func ReadFrom[T_from any](buffer *Buffer, from io.Reader[T_from]) (n int64, err error) {
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
	r, err = from.Read(from.Object, slice[l:size])

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
