package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count uint
	mutex sync.Mutex
}

func (c *Counter) Add() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func workCount(c *Counter, wg *sync.WaitGroup) {
	for i := 0; i < 50; i++ {
		c.Add()
	}
	wg.Done()
}

func main() {
	const workersCount = 5

	var wg sync.WaitGroup
	wg.Add(workersCount)

	var counter Counter

	for i := 0; i < workersCount; i++ {
		go workCount(&counter, &wg)
	}
	wg.Wait()

	fmt.Println("Counter: ", counter.count)
}
