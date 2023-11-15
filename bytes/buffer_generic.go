//go:build !amd64

package bytes

func Reset(buffer *Buffer) {
	*buffer = (*buffer)[:0]
}

func (buffer *Buffer) Reset() {
	Reset(buffer)
}

func String(buffer *Buffer) string {
	return string(*buffer)
}

func (buffer *Buffer) String() string {
	return String(buffer)
}
