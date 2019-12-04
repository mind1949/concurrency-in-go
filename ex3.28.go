// page: 81
// select语句
// 使用time.After避免永久阻塞
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Time out")
	}
}
