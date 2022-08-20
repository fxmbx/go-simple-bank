package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//alphabets that contains all acceptable
const alphabets = "abcdefghijklmonpqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // return a random integer between 0- (max-min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

//Generate a random owner name
func RandomOwner() string {
	return RandomString(6)
}

//generate a random int fo amount
func RandomMoney() int64 {
	return RandomInt(0, 100)
}

//Rancom currency selects a random currency code
func RandomCurrency() string {
	currency := []string{CAD, NGN}
	n := len(currency)
	return currency[rand.Intn(n)]
}

func RandomAccountId() int64 {
	return RandomInt(1, 39)
}

func AccountId() int64 {
	return 1
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(8))
}
