package utils

import (
	"math/rand"
	"time"
)

func RandomInt(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max) + 1
}
