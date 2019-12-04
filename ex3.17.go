// page: 66
// chan
package main

// go的类型系统不允许我们对只读通道做写操作,
// 也不允许对只写通道执行读操作
func main() {
	writeStream := make(chan<- interface{})
	readStream := make(<-chan interface{})

	<-writeStream
	readStream <- struct{}{}
}

// output:
//# command-line-arguments
//./ex3.17.go:7:2: invalid operation: <-writeStream (receive from send-only type chan<- interface {})
//./ex3.17.go:8:13: invalid operation: readStream <- struct {} literal (send to receive-only type <-chan interface {})
