// page: 58
// Once
package main

import (
	"fmt"
	"sync"
)

// 最终会输出Count: 1
// 而不是Count: 0
// 这是因为Do记录的是自身被执行的次数,
// 而不是Do传入的函数被执行的次数
func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

// output:
// Count: 1
