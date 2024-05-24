package percent

import (
	"github.com/gohryt/asphyxia/bytes"
)

const (
	upperhex = "0123456789ABCDEF"
	lowerhex = "0123456789abcdef"
)

func Encode(source bytes.Buffer) bytes.Buffer {
	var (
		i        int
		j        int
		capacity int
		b        byte
	)

	for i < len(source) {
		b = source[i]

		if QuotedPathShouldEscapeTable[int(b)] != 0 {
			capacity += 3
		} else {
			capacity += 1
		}

		i += 1
	}

	target := make(bytes.Buffer, capacity)

	if capacity > len(source) {
		i = 0

		for i < len(source) {
			b = source[i]

			if QuotedPathShouldEscapeTable[int(b)] != 0 {
				target[j+2] = upperhex[b&0xf]
				target[j+1] = upperhex[b>>4]
				target[j] = '%'

				j += 3
			} else {
				target[j] = b

				j += 1
			}

			i += 1
		}

		target = target[:j]
	} else {
		copy(target, source)
	}

	return target
}

func Decode(source bytes.Buffer) bytes.Buffer {
	var (
		i        int
		j        int
		capacity int
		b        byte
	)

	for i < len(source) {
		capacity += 1

		if source[i] == '%' {
			i += 3
		} else {
			i += 1
		}
	}

	target := make(bytes.Buffer, capacity)

	if i > capacity {
		i = 0
		capacity = 0

		for i < len(source) {
			b = source[i]

			if b == '%' {
				j = i + 2

				if j >= len(source) {
					copy(target[capacity:], source[i:])
					return target
				}

				x2 := Hex2IntTable[source[j]]
				x1 := Hex2IntTable[source[i+1]]

				if x1 == 16 || x2 == 16 {
					target[capacity] = '%'
				} else {
					target[capacity] = x1<<4 | x2
					i = j
				}
			} else {
				target[capacity] = b
			}

			i += 1
			capacity += 1
		}
	} else {
		copy(target, source)
	}

	return target
}
