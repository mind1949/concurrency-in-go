// page: 42
// 解决捕获迭代变量问题
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // 通过传入一个copy来解决捕获迭代变量问题
	}
	wg.Wait()
}
