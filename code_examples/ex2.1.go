// page: 40
// fork-join model
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello() // fork child goroutine
	wg.Wait()     // join child goroutine
}
