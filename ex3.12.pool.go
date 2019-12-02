// page: 60
// sync.Pool
package main

import (
	"fmt"
	"sync"
)

// 正确使用sync.Pool的Get与Put函数后,
// 可以大大减少初始化实例的数量
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	for i := 0; i < 100; i++ {
		instance := myPool.Get()
		myPool.Put(instance)
	}
}

// output:
//Creating new instance.
//Creating new instance.
