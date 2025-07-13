package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func work(dataCh chan int, ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("worker", <-dataCh, "start")
	for data := range dataCh {
		select {
		case <-ctx.Done():
			fmt.Println("worker", <-dataCh, "closed")
			wg.Done()
			return
		default:
			fmt.Fprintln(os.Stdout, data)
			time.Sleep(500 * time.Millisecond)
		}
	}
	wg.Done()
}

func main() {
	const workersCount = 20

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(workersCount)

	dataCh := make(chan int, workersCount)
	exitCh := make(chan os.Signal, 1)

	//Workers pool
	for i := 0; i < workersCount; i++ {
		go work(dataCh, ctx, &wg)
	}

	//Shutdown
	go func() {
		signal.Notify(exitCh, syscall.SIGINT, syscall.SIGTERM)
		if <-exitCh != nil {
			fmt.Println("\nReceived shutdown signal, initiating shutdown...")
			close(exitCh)
			cancel()
			wg.Wait()
			close(dataCh)
			os.Exit(0)
		}
	}()

	//Сontinuous record
	for {
		dataCh <- rand.Intn(100)
	}

}
