package bytes

import (
	"io"
	"unicode/utf8"

	"github.com/gohryt/asphyxia/memory"
)

type (
	Buffer struct {
		Data []byte
	}
)

func BufferFrom(value []byte) *Buffer {
	return &Buffer{
		Data: value,
	}
}

func BufferFromString(value string) *Buffer {
	return &Buffer{
		Data: []byte(value),
	}
}

func (buffer *Buffer) String() string {
	return string(buffer.Data)
}

func (buffer *Buffer) Reset() {
	buffer.Data = buffer.Data[:0]
}

func (buffer *Buffer) Clone() *Buffer {
	return &Buffer{
		Data: append([]byte(nil), buffer.Data...),
	}
}

func (buffer *Buffer) Set(value []byte) {
	buffer.Data = append(buffer.Data[:0], value...)
}

func (buffer *Buffer) SetString(value string) {
	buffer.Data = append(buffer.Data[:0], value...)
}

func (buffer *Buffer) Write(value []byte) (n int, err error) {
	buffer.Data = append(buffer.Data, value...)
	return len(value), nil
}

func (buffer *Buffer) WriteString(value string) (n int, err error) {
	buffer.Data = append(buffer.Data, value...)
	return len(value), nil
}

func (buffer *Buffer) WriteByte(value byte) (err error) {
	buffer.Data = append(buffer.Data, value)
	return nil
}

func (buffer *Buffer) WriteRune(value rune) (n int, err error) {
	slice := buffer.Data
	l := len(slice)

	size := l + utf8.UTFMax

	if size > cap(slice) {
		reallocation := make([]byte, size)
		copy(reallocation, slice)

		slice = reallocation
	}

	n = utf8.EncodeRune(slice[l:size], value)

	buffer.Data = slice[:(l + n)]
	return n, nil
}

func (buffer *Buffer) ReadFrom(from io.Reader) (n int64, err error) {
	slice := buffer.Data
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

	reallocation := make([]byte, size)
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

	buffer.Data = slice[:l]
	return n, err
}
