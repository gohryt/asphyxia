//go:build amd64

package bytes

func reset(buffer *Buffer)

func (buffer *Buffer) Reset() {
	reset(buffer)
}

func asString(*Buffer) string

func (buffer *Buffer) AsString() string {
	return asString(buffer)
}
