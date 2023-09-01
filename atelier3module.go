package atelier3module

import (
	"math/rand"
)

func ArrayGenerate(size int) []int {
	var slice []int = make([]int, size)

	for index := range slice {
		slice[index] = rand.Intn(10000) + 1
	}

	return slice
}
