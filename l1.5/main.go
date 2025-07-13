package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const runTimeS = 10
	const tickRateMs = 500

	dataCh := make(chan int, 5)

	timeout := time.After(runTimeS * time.Second)
	ticker := time.NewTicker(tickRateMs * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			close(dataCh)
			fmt.Println("Exit")
			return

		case <-ticker.C:
			num := rand.Intn(100)
			select {
			case dataCh <- num:
				fmt.Println("Send:", num)
			default:
				fmt.Println("Channel blocked")
			}

		case received := <-dataCh:
			fmt.Println("Received:", received)
		}
	}
}
