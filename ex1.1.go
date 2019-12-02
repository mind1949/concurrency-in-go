package main

import "fmt"

// 数据竞态: page 4
func main() {
	var data int
	go func() {
		data++ // ①
	}()
	if data == 0 { // ②
		fmt.Printf("the value is %v.\n", data) // ③
	}
}
