package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "abcd"
	counts := make(map[string]int)

	lowerText := strings.ToLower(text)

	for _, s := range lowerText {
		_, ok := counts[string(s)]
		if ok {
			fmt.Println("false")
			return
		} else {
			counts[string(s)] = 1
		}
	}
	fmt.Println("true")
}
