// page: 45
// 测量goroutine上下文切换耗费的时间
package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

// go test -bench=. -cpu=1 \
//> code_examples/ex2.5_test.go
//goos: darwin
//goarch: amd64
//BenchmarkContextSwitch  10000000               167 ns/op
//PASS
//ok      command-line-arguments  1.856s
