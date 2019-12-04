// page: 68
// chan
// 读取通道时, 可选择性的读取
// 第二个值(表示chan是否被close)
package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello World!"
	}()
	// 关闭前
	salutation, ok := <-stringStream
	fmt.Printf("被关闭(%v): %v\n", !ok, salutation)
	// 关闭后
	close(stringStream)
	salutation, ok = <-stringStream
	fmt.Printf("被关闭(%v): %v\n", !ok, salutation)
}
