package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun moon lunar mars"
	words := strings.Fields(text)

	for i := 0; i < len(words)/2; i++ {
		reverseIndex := len(words) - i - 1
		words[i], words[reverseIndex] = words[reverseIndex], words[i]
	}

	fmt.Println(words)
}
