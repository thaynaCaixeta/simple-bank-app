package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "ashjgfddbfjewhfiwbhfwkf"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63()
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func GenerateRandomOwner() string {
	return RandomString(6)
}

func GenerateRandomMoney() int64 {
	return RandomInt(0, 100000)
}

func GenerateRandomCurrency() string {
	currencies := []string{"EUR", "USD", "BLR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
