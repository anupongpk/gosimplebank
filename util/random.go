package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphaet = "abcdefghijklmnopqrstuvwxyz"

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random number
func RandomInt(min, max int64) int64 {
	return min + rnd.Int63n(max-min+1)
}

// RandomString generates a random string of lenth n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphaet)

	for i := 0; i < n; i++ {
		c := alphaet[rnd.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency cod
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rnd.Intn(n)]
}
