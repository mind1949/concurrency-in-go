package main

import (
	"fmt"
	"time"
)

// 虽然通过休眠能够使执行结果正确,
// 但是逻辑上并不正确
func main() {
	var data int
	go func() { data++ }()
	time.Sleep(1 * time.Second) // this is bad!
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
