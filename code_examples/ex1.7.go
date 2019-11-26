// page: 16
// 饥饿
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loop\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			for i := 0; i < 3; i++ {
				sharedLock.Lock()
				time.Sleep(1 * time.Nanosecond)
				sharedLock.Unlock()
			}

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loop\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}

// output:
//	Polite worker was able to execute 395403 work loop
//	Greedy worker was able to execute 799401 work loop
