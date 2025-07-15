package main

import (
	"fmt"
	"time"
)

func sleep(ns int64) {
	timer := time.Now()

	for time.Since(timer).Milliseconds() < ns {
	}
}

func main() {
	start := time.Now()
	fmt.Println("Start")

	sleep(10000)

	fmt.Println("Time:", time.Since(start))
}
