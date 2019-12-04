// page: 82
// for-select loop
// 用于让goroutine一边做某些工作,
// 一遍等待另一个goroutine发送消息
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

// output:
// Achieved 5 cycles of work before signalled to stop.
