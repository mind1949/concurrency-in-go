// page: 89
// for-select loop
// 结果可能会输出:
// a
//close(done)
//b
// 是正常的,
// 但是居然还有可能输出:
// a
//close(done)
//b
//c
// 这就很奇怪了,
// 我认为原因是这样:
// 1. 一定是for range channel每从channel
// 中读取一个数值就执行一次循环体中的代码,
// 当循环体中的代码被执行完了, 才会再次读取channel中的数值.
// 2. 之所以会父goroutine再读取到b便关闭done通道的时候依然能够打印出c
// 是因为有时候关闭done的时候, 子goroutine已经进行到了下一个循环,
// 并且已经到了①以下
package main

import "fmt"

func main() {
	forSelect := func(done <-chan struct{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for _, s := range []string{"a", "b", "c", "d", "e"} {
				select {
				case <-done: // ①
					return
				case stringStream <- s:
				}
			}
		}()
		return stringStream
	}

	done := make(chan struct{})
	stringStream := forSelect(done)
	for s := range stringStream {
		if s == "b" {
			fmt.Println("close(done)")
			close(done)
		}
		fmt.Println(s)
	}
}

// output:
// a
//close(done)
//b
//c
// 或者
//a
//close(done)
//b
