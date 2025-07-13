package main

import (
	"fmt"
)

func main() {
	text := "главрыба"
	symbols := []rune(text)

	for i := 0; i < len(symbols)/2; i++ {
		reverseIndex := len(symbols) - i - 1
		symbols[i], symbols[reverseIndex] = symbols[reverseIndex], symbols[i]
	}

	fmt.Println(string(symbols))
}
