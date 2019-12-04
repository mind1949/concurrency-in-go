// page:67
// chan
// 通道阻塞,造成死锁!
package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		stringStream <- "hello world"
	}()

	fmt.Println(<-stringStream)
}
