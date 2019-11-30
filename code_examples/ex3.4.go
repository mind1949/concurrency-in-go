// page: 49
// 互斥锁
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrement: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete")
}

// output:
//Incrementing: 1
//Decrement: 0
//Incrementing: 1
//Incrementing: 2
//Incrementing: 3
//Incrementing: 4
//Decrement: 3
//Decrement: 2
//Decrement: 1
//Incrementing: 2
//Decrement: 1
//Decrement: 0
//Arithmetic complete
