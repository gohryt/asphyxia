package bytes

import (
	"io"
	"unicode/utf8"

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

func Set(buffer *Buffer, source []byte) {
	*buffer = append((*buffer)[:0], source...)
}

func SetString(buffer *Buffer, source string) {
	*buffer = append((*buffer)[:0], source...)
}

func Write(buffer *Buffer, source []byte) (n int, err error) {
	*buffer = append(*buffer, source...)
	return len(source), nil
}

func WriteByte(buffer *Buffer, source byte) (err error) {
	*buffer = append(*buffer, source)
	return
}

func WriteRune(buffer *Buffer, source rune) (n int, err error) {
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

func WriteString(buffer *Buffer, source string) (n int, err error) {
	*buffer = append(*buffer, source...)
	return len(source), nil
}

func ReadFrom(buffer *Buffer, source io.Reader) (n int64, err error) {
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
