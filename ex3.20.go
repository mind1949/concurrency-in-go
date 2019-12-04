// page: 68
// chan
// 通过close来关闭chan.
// 1. 被关闭的可读通道允许被读取无数次,
// 以保证上游的通道被关闭后, 下游的进程能够正常读取信息
// 2. 从被关闭的通道中读物的值是通道类型的零值(zero value)

package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer)
}
