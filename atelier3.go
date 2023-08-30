package atelier3

import (
	"math/rand"
)

func Array_Generate(x int) (s []int) {
	for i := 0; i < x; i++ {
		s = append(s, rand.Intn(10000))
	}
	return
}
