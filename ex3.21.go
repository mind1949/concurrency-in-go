// page: 69
// chan
// close提供了一个使用for和range
// 来优雅的遍历通道的方式.
// 0. 退出条件是通道被关闭;
// 1. 不返回第二个布尔值;
// 2. 不返回通道被关闭是返回的默认值;
package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
	fmt.Println("")
}

// output:
// 1 2 3 4 5
