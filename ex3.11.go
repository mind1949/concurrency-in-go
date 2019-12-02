// page: 59
// sync.Once
package main

import (
	"fmt"
	"sync"
)

// 与ex3.10中sync.Once避免多次执行的方式不同,
// 在这部分代码中sync.Once会通过死锁来避免循环引用
func main() {
	var onceA, onceB sync.Once
	var initB func()

	initA := func() {
		fmt.Printf("执行到: initA\n")
		onceB.Do(initB)
	}
	initB = func() {
		fmt.Printf("执行到: initB\n")
		onceA.Do(initA)
	}

	onceA.Do(initA)
}

//output:
//执行到: initA
//执行到: initB
//fatal error: all goroutines are asleep - deadlock!
