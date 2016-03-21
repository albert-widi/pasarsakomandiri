package tools

import (
	"math/rand"
	"sync"
)

var runes  =  []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var mux sync.Mutex

func RandomString(l int) string {
	runesLen := len(runes)
	bytes := make([]rune, l)
	for i:=0 ; i<l ; i++ {
		mux.Lock()
		bytes[i] = runes[rand.Intn(runesLen)]
		mux.Unlock()
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}