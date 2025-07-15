package main

import (
	"math/rand"
)

var justString string

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Possible implementation
func createHugeString(size int) []rune {
	newStr := make([]rune, size)
	for i := range newStr {
		newStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return newStr
}

func someFunc() {
	size := 1 << 10
	v := createHugeString(size)
	if size < 100 {
		justString = string(v[:size])
		return
	}
	justString = string(v[:100])
}

func main() {
	someFunc()

	println(justString)
}
