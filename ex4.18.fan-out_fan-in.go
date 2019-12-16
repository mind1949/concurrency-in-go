// page: 114
// fan-out, fan-in
// 扇出, 扇入
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand := func() interface{} { return rand.Intn(5e7) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
