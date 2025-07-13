package main

import (
	"fmt"
)

func getUniques(strings []string) (uniques []string) {
	uniquesMap := make(map[string]any) //any have size 0

	for _, str := range strings {
		_, ok := uniquesMap[str]
		if !ok {
			uniquesMap[str] = struct{}{}
			uniques = append(uniques, str)
		}
	}
	return uniques

}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	uniques := getUniques(strings)
	fmt.Println(uniques)
}
