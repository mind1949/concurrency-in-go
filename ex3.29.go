// page: 81
// select statement
// 使用default语句保证,
// 没有一个case语句满足条件时,
// 直接执行default语句下的内容
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c1, c2 := make(<-chan int), make(<-chan int)
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
