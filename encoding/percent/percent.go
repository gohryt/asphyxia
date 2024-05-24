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
	var (
		i        int
		capacity int

		r    rune
		size int
		u    uint32
	)

	for i < len(source) {
		r, size = UTF8.DecodeRune(source[i:])
		u = uint32(r)

		i += size

		switch {
		case u <= UTF8.Rune1Max:
			capacity += 3
		case u <= UTF8.Rune2Max:
			capacity += 6
		case u <= UTF8.Rune3Max:
			capacity += 9
		default:
			capacity += 12
		}
	}

	target := make(bytes.Buffer, 0, capacity)

	for _, c := range source {
		if QuotedPathShouldEscapeTable[int(c)] != 0 {
			target = append(target, '%', upperhex[c>>4], upperhex[c&0xf])
		} else {
			target = append(target, c)
		}
	}

	return target
}

func Decode(source bytes.Buffer) bytes.Buffer {
	var (
		i        int
		capacity int
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

		var (
			b byte
			j int
		)

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
