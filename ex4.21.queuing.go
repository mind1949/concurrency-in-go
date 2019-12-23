package main

import (
	"fmt"
	"time"
)

func main() {
	repeat := func(
		done <-chan interface{},
		elements ...interface{},
	) <-chan interface{} {
		dataStream := make(chan interface{})
		go func() {
			defer close(dataStream)
			for {
				for _, e := range elements {
					select {
					case <-done:
						return
					case dataStream <- e:
					}
				}
			}
		}()
		return dataStream
	}

	take := func(
		done <-chan interface{},
		count int,
		inputStream <-chan interface{},
	) <-chan interface{} {
		dataStream := make(chan interface{})
		go func() {
			defer close(dataStream)
			for i := 0; i < count; i++ {
				select {
				case <-done:
					return
				case dataStream <- <-inputStream:
				}
			}
		}()
		return dataStream
	}

	sleep := func(
		done <-chan interface{},
		duration time.Duration,
		inputStream <-chan interface{},
	) <-chan interface{} {
		dataStream := make(chan interface{})
		go func() {
			defer close(dataStream)
			for e := range inputStream {
				select {
				case <-done:
					return
				case <-time.After(duration):
					fmt.Printf("sleep: %v\n", duration)
					dataStream <- e
				}
			}
		}()
		return dataStream
	}

	done := make(chan interface{})
	defer close(done)

	zeros := take(done, 3, repeat(done, 0))
	short := sleep(done, 1*time.Second, zeros)
	long := sleep(done, 4*time.Second, short)
	pipeline := long

	now := time.Now()
	for e := range pipeline {
		fmt.Println(e)
	}
	fmt.Printf("elapsed: %v\n", time.Since(now))
}
