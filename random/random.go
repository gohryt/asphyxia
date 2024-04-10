package random

import (
	"math/rand/v2"
)

func Slice[T any](source []T, length int) []T {
	max := len(source)

	target := make([]T, length)

	for i := range target {
		target[i] = source[rand.IntN(max)]
	}

	return target
}
