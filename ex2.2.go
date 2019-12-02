// page: 41
// 捕获迭代变量
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation) // 会捕获迭代变量, 导致逻辑错误
		}()
	}
	wg.Wait()
}

// output:
// good day
// good day
// good day
