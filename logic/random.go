package logic

import (
	"math/rand"
)

func randomPos() (int, int) {
	c := rand.Intn(4)
	r := rand.Intn(4)
	return c, r
}

func randomValue() int {
	return 2 + 2*rand.Intn(2)
}
