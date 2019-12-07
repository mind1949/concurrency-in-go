// page: 93
// 预防goroutine泄露
// 从输出的结果可以看到没有打印出"newRandStream closure exited.",
// 所以可以判断newRandStream孵化的goroutine会内存泄露
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				rand.Seed(time.Now().UnixNano())
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

// output:
// 3 random ints:
//0: 1033292541545608891
//1: 9169269580881113580
//2: 1409402647487455247
//
