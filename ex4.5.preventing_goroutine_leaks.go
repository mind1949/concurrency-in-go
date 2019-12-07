// page: 91
// 预防go协成泄露

// 在本例中因为doWork是在主goroutine中启动的,
// 所以会导致死锁.
// 但若doWork是在某个子goroutine中
// 启动就会导致goroutine泄露
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer wg.Done()
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	wg.Add(1)
	doWork(nil)
	wg.Wait()
	// Perhaps more work is done
	fmt.Println("Done.")
}
