package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomInt(max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(a.Int64())
}

func RandomIntBetween(min int, max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(a.Int64()) + min
}
