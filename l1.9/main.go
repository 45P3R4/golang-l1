package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func producer(numbers [10]int, numCh chan int) {
	for _, num := range numbers {
		numCh <- num
		time.Sleep(500 * time.Millisecond)
	}
	close(numCh)
}

func square(numCh chan int, sqrCh chan int) {
	for num := range numCh {
		sqrCh <- num * 2
	}
	close(sqrCh)
}

func reciever(sqrCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for s := range sqrCh {
		fmt.Fprintln(os.Stdout, s)
	}
}

func main() {
	numbers := [10]int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	numCh := make(chan int)
	sqrCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go producer(numbers, numCh)
	go square(numCh, sqrCh)
	go reciever(sqrCh, &wg)

	wg.Wait()
}
