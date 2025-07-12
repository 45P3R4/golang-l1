package main

import (
	"fmt"
	"math/rand"
	"os"
)

func work(dataCh chan int) {
	for data := range dataCh {
		fmt.Fprintln(os.Stdout, data)
	}
}

func main() {
	const workersCount = 4

	dataCh := make(chan int, 5)

	for i := 0; i < workersCount; i++ {
		go work(dataCh)
	}

	for {
		num := rand.Intn(100)
		dataCh <- num
	}

}
