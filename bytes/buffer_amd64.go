//go:build amd64

package bytes

func Reset(buffer *Buffer)

func (buffer *Buffer) Reset() {
	Reset(buffer)
}

func String(*Buffer) string

func (buffer *Buffer) String() string {
	return String(buffer)
}
