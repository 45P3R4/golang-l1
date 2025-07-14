package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func stopMeWithNone(wg *sync.WaitGroup) {
	defer fmt.Println("No thanks, I can stop myself")
	defer wg.Done()
}

func stopMeWithChannel(stopCh chan any, wg *sync.WaitGroup) {
	defer fmt.Println("Stopped by channel")
	defer wg.Done()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-stopCh:
			return
		case <-ticker.C:
			fmt.Println("Stop me")
		}
	}
}

func stopMeWithTimeout(wg *sync.WaitGroup) {
	defer fmt.Println("Stopped by timeout")
	defer wg.Done()
	ticker := time.NewTicker(500 * time.Millisecond)
	timeout := time.After(time.Duration(rand.Intn(5)) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-timeout:
			return
		case <-ticker.C:
			fmt.Println("Stop me")
		}
	}
}

func stopMeWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer fmt.Println("Stopped by context")
	defer wg.Done()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("Stop me")
		}
	}
}

func stopMeWithGoexit(wg *sync.WaitGroup) {
	defer fmt.Println("Stopped by goexit")
	defer wg.Done()
	fmt.Println("Stop me")
	runtime.Goexit()
}

func main() {
	stopCh := make(chan any)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(5)

	go stopMeWithNone(&wg)

	go stopMeWithChannel(stopCh, &wg)

	go stopMeWithTimeout(&wg)

	go stopMeWithContext(ctx, &wg)

	go stopMeWithGoexit(&wg)

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	stopCh <- struct{}{}
	cancel()

	wg.Wait()

}
