package random

import (
	"math/rand"
	"time"
)

var GlobalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Slice[T any](source []T, length int) []T {
	max := len(source)

	target := make([]T, length)

	for i := range target {
		target[i] = source[GlobalRand.Intn(max)]
	}

	return target
}
