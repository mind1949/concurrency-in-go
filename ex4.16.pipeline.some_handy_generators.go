package main

import "fmt"

func main() {
	// repeat Generator
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	// take
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	// add type assertion stage
	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	// combine stages
	done := make(chan interface{})
	defer close(done)

	//for num := range take(done, repeat(done, 1), 10) {
	//	fmt.Printf("%v ", num)
	//}
	var message string
	for token := range toString(
		done,
		take(done, repeat(done, "I", "Love", "You", "!"), 16),
	) {
		message += token
	}
	fmt.Printf("message: %s...\n", message)

}
