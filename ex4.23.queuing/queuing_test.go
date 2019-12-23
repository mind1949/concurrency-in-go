package queuing

import (
	"bufio"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func repeat(
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

func take(
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

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B) {
	bufferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferedFile))
}

func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, b.N, repeat(done, byte(0))) {
		writer.Write([]byte{bt.(byte)})
	}
}
