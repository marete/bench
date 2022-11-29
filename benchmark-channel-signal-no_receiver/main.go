package main

import (
	"fmt"
	"time"
)

const numIterations = 2 << 20

func main() {
	ch := make(chan struct{})
	start := time.Now()
	for i := 0; i < numIterations; i++ {
		select {
		case ch <- struct{}{}:
			panic("unexpected")
		default:
			continue
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%d µs elapsed -- %7.3f µs per operation\n", elapsed.Microseconds(),
		float64(elapsed.Microseconds())/float64(numIterations))
}
