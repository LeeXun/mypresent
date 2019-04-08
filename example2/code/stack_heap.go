package main

import (
	"fmt"
	"time"
)

var N int = 1 << 15

func BenchmarkAllocateFooStack() {
	for i := 0; i < N; i++ {
		var a [1 << 18]byte
		a[0] = 1
	}
}

func BenchmarkAllocateFooHeap() {
	for i := 0; i < N; i++ {
		var bts = make([]byte, 1<<18)
		bts[0] = 1
	}
}

func main() {
	t0 := time.Now()
	BenchmarkAllocateFooStack()
	t1 := time.Now()
	BenchmarkAllocateFooHeap()
	t2 := time.Now()
	fmt.Printf("%v\n", t1.Sub(t0))
	fmt.Printf("%v\n", t2.Sub(t1))
}

// go get -u -v golang.org/x/tools/cmd/present
