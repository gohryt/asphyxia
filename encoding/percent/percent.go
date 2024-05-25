package percent

import (
	"github.com/gohryt/asphyxia/bytes"
	"github.com/gohryt/asphyxia/encoding/hex"
)

func Encode(source *bytes.Buffer) *bytes.Buffer {
	var (
		i        int
		j        int
		capacity int
		b        byte
	)

	sourceData := source.Data

	for i < len(sourceData) {
		b = sourceData[i]

		if ShouldEscapeTable[int(b)] != 0 {
			capacity += 3
		} else {
			capacity += 1
		}

		i += 1
	}

	if capacity == len(sourceData) {
		return source.Clone()
	}

	targetData := make([]byte, capacity)

	i = 0

	for i < len(sourceData) {
		b = sourceData[i]

		if ShouldEscapeTable[int(b)] != 0 {
			targetData[j+2] = hex.Upper[b&0xf]
			targetData[j+1] = hex.Upper[b>>4]
			targetData[j] = '%'

			j += 3
		} else {
			targetData[j] = b

			j += 1
		}

		i += 1
	}

	return &bytes.Buffer{
		Data: targetData,
	}
}

func Decode(source *bytes.Buffer) *bytes.Buffer {
	var (
		i        int
		j        int
		capacity int
		b        byte

		x2 byte
		x1 byte
	)

	sourceData := source.Data

	for i < len(sourceData) {
		capacity += 1

		if sourceData[i] == '%' {
			i += 3
		} else {
			i += 1
		}
	}

	if i == capacity {
		return source.Clone()
	}

	targetData := make([]byte, capacity)

	i = 0
	capacity = 0

	for i < len(sourceData) {
		b = sourceData[i]

		if b == '%' {
			j = i + 2

			if j >= len(sourceData) {
				copy(targetData[capacity:], sourceData[i:])
				return &bytes.Buffer{
					Data: targetData,
				}
			}

			x2 = hex.HexToIntTable[sourceData[j]]
			x1 = hex.HexToIntTable[sourceData[i+1]]

			if x1 == 16 || x2 == 16 {
				targetData[capacity] = '%'
			} else {
				targetData[capacity] = x1<<4 | x2
				i = j
			}
		} else {
			if b == '+' {
				targetData[capacity] = ' '
			} else {
				targetData[capacity] = b
			}
		}

		i += 1
		capacity += 1
	}

	return &bytes.Buffer{
		Data: targetData,
	}
}
