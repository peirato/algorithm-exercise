package number_util

import (
	"crypto/rand"
	"math/big"
)

func randomInt(max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(a.Int64())
}

func randomIntBetween(min int, max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(a.Int64()) + min
}
