package main

import (
	"math/rand"
)

var justString string

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Possible implementation
func createHugeString(size int) string {
	newStr := make([]rune, size)
	for i := range newStr {
		newStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(newStr)
}

func someFunc() {
	size := 1 << 10
	v := createHugeString(size)
	if size < 100 {
		justString = string(v[:size])
		return
	}
	justString = v[:100]
}

func main() {
	someFunc()

	println(justString)
}
