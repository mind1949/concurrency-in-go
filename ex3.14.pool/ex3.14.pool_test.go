package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v\n", err)
		}
		if _, err = ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v\n", err)
		}
		conn.Close()
	}
}

// output:
//goos: darwin
//goarch: amd64
//BenchmarkNetworkRequest-8             10        1004527934 ns/op
//PASS
//ok      _/Users/mind1949/roadmap/golang/concurrency_in_go/ex3.14.pool   11.060s
