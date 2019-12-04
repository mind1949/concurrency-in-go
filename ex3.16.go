// page: 65
// chan
// ## 声明通道与实例化通道
// ### 1. 声明以个通道
// `var dataStream chan interface{}`
// ### 2. 初始化化一个通道
// `dataStream := make(chan interface{})`
// ### 3. 声明一个单向通道
// `var dataStream <- chan interface{}`
// ### 4. 初始化一个单向通道
// `dataStream := make(chan interface{})`
// ### 5. 声明与实例化只写或只读通道
// ```go
// var receiveChan <- chan interface{}
// var sendChan chan interface{} <-
// receiveChan = dataStream
// sendStream = dataStream
// ```

package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channels!"
	}()
	fmt.Println(<-stringStream)
}
