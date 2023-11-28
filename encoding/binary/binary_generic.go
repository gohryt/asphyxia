//go:build !amd64

package binary

func BigEndianCopyUint16(to []byte, from ...uint16) int {
	i := 0
	j := 0

	count := len(to) / 2

	if len(from) < count {
		count = len(from)
	}

	for i < count {
		element := from[i]

		to[j] = byte(element >> 8)
		to[j+1] = byte(element)

		i += 1
		j += 2
	}

	return j
}
