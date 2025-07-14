package main

import (
	"math/rand"
	"sync"
)

type SafeMap struct {
	data  map[int]int
	mutex sync.RWMutex
}

func (sm *SafeMap) New() {
	sm.data = make(map[int]int)
}

func (sm *SafeMap) Set(key int, value int) {
	sm.mutex.Lock()
	sm.data[key] = value
	sm.mutex.Unlock()
}

func writeWorker(reciever map[int]int, mutex *sync.Mutex) {
	for {
		index := rand.Intn(100)
		mutex.Lock()
		reciever[index] = index
		mutex.Unlock()
	}
}

func safeMapWriteWorker(reciever *SafeMap) {
	for {
		index := rand.Intn(100)
		reciever.Set(index, index)
	}
}

func main() {
	const workersCount = 3
	reciever := make(map[int]int)

	var safeMap SafeMap
	safeMap.New()

	var mutex sync.Mutex

	//Mutex in worker
	for i := 0; i < workersCount; i++ {
		go writeWorker(reciever, &mutex)
	}

	//Mutex in structure type
	for i := 0; i < workersCount; i++ {
		go safeMapWriteWorker(&safeMap)
	}
}
