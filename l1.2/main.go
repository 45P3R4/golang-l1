package main

import (
	"fmt"
	"os"
	"strconv"
)

func printSquare(number int) {
	fmt.Fprintf(os.Stdout, strconv.Itoa(number*number)+"\n")
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	for _, num := range numbers {
		go printSquare(num)
	}

	fmt.Scanln()
}
