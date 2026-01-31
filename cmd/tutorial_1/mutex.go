package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
// shared
var counter int

func increment(wg *sync.WaitGroup) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// all 1000 increments complete before printing
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

// Worker A enters → increments → leaves
// Worker B enters → increments → leaves
// Worker C enters → increments → leaves