package binary

func BigEndianCopyUint16(to []byte, from ...uint16) int {
	i := 0
	j := 0

	count := min(len(from), len(to)/2)

	for i < count {
		element := from[i]

		to[j] = byte(element >> 8)
		to[j+1] = byte(element)

		i += 1
		j += 2
	}

	return j
}
