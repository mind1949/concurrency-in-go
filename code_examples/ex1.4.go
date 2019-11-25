package main

import (
	"fmt"
	"sync"
)

// 使用同步锁能够解决数据竞态问题,
// 但是无法解决竞态条件问题
func main() {
	var memoryAccess sync.Mutex
	var data int
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}
