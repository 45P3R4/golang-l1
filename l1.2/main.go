package main

import (
	"fmt"
	"os"
	"sync"
)

func printSquare(number int, wg *sync.WaitGroup) {
	fmt.Fprintf(os.Stdout, "Square of %d is %d\n", number, number*number)
	wg.Done()
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, num := range numbers {
		go printSquare(num, &wg)
	}

	wg.Wait()
}
