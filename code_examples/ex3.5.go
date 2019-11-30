package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Reader\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

// output:
//Reader  RWMutext      Mutex
//1       56.861µs      4.132µs
//2       15.76µs       34.606µs
//4       8.093µs       3.826µs
//8       31.31µs       16.105µs
//16      34.507µs      45.286µs
//32      43.857µs      69.409µs
//64      124.787µs     109.976µs
//128     96.229µs      164.342µs
//256     228.56µs      148.277µs
//512     333.625µs     233.305µs
//1024    354.466µs     528.175µs
//2048    700.604µs     852.908µs
//4096    1.353213ms    1.727831ms
//8192    2.7778ms      3.840098ms
