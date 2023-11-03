//go:build !amd64

package bytes

func (buffer *Buffer) Reset() {
	*buffer = (*buffer)[:0]
}

func (buffer *Buffer) AsString() string {
	return string(*buffer)
}
