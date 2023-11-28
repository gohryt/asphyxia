package percent

import (
	"github.com/gohryt/asphyxia/bytes"
	"github.com/gohryt/asphyxia/unicode/UTF8"
)

const (
	upperhex = "0123456789ABCDEF"
	lowerhex = "0123456789abcdef"
)

func Encode(source bytes.Buffer) bytes.Buffer {
	i := 0
	c := 0

	for i < len(source) {
		rune, size := UTF8.DecodeRune(source[i:])
		uinted := uint32(rune)

		i += size

		switch {
		case uinted <= UTF8.Rune1Max:
			c += 3
		case uinted <= UTF8.Rune2Max:
			c += 6
		case uinted <= UTF8.Rune3Max:
			c += 9
		default:
			c += 12
		}
	}

	target := make(bytes.Buffer, c)

	if c > len(source) {
		i = len(source) - 1
		c -= 1

		for i >= 0 {
			byte := source[i]

			i -= 1

			if quotedPathShouldEscapeTable[byte] != 0 {
				target[c] = upperhex[byte&0xf]
				target[c-1] = upperhex[byte>>4]
				target[c-2] = '%'
				c -= 3
			} else {
				target[c] = byte
				c -= 1
			}
		}
	} else {
		copy(target, source)
	}

	return target
}

func Decode(source bytes.Buffer) bytes.Buffer {
	c := 0
	i := 0

	for i < len(source) {
		c += 1

		if source[i] == '%' {
			i += 3
		} else {
			i += 1
		}
	}

	target := make(bytes.Buffer, c)

	if i > c {
		i = 0
		c = 0

		for i < len(source) {
			byte := source[i]

			if byte == '%' {
				j := i + 2

				if j >= len(source) {
					copy(target[c:], source[i:])
					return target
				}

				x2 := hex2intTable[source[j]]
				x1 := hex2intTable[source[i+1]]

				if x1 == 16 || x2 == 16 {
					target[c] = '%'
				} else {
					target[c] = x1<<4 | x2
					i = j
				}
			} else {
				target[c] = byte
			}

			i += 1
			c += 1
		}
	} else {
		copy(target, source)
	}

	return target
}
