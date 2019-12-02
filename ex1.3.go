package main

import (
	"fmt"
)

// 两个进程对同一个变量执行读写操作,
// 因为读写顺序不确定, 造成了数据竞态
func main() {
	var data int
	go func() { data++ }()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}
