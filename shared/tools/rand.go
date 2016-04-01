package tools

import (
	"math/rand"
)

var runes  =  []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(l int) string {
	runesLen := len(runes)
	bytes := make([]rune, l)
	for i:=0 ; i<l ; i++ {
		bytes[i] = runes[rand.Intn(runesLen)]
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}