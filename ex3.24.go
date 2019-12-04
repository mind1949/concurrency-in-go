// page: 77
// channel
// 通道的所有权
package main

import "fmt"

func main() {
	// 通道的拥有者负责实例化/写/关闭通道
	// 尽量保证通道所有者的词法作用域尽量小
	// 以保证能够清晰的看到通道的实例化/写/关闭
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	// 通道的消费者负责
	// 读/思考通道阻塞或者关闭
	// 时如何处理;
	// 通道消费者只接受的通道应该为只读通道;
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
