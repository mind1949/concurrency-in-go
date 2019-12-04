// page: 82
// block forever
// 永久阻塞, 死锁!
package main

func main() {
	select {}
}
