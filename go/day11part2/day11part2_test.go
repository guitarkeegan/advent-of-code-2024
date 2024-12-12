package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkBlink(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := strings.Split(input, " ")
		s := stones(Map(data, atoi))
		dbg("s length: %d", len(s))
		for i := 0; i < 25; i++ {
			res, err := blink(s)
			if err != nil {
				fmt.Errorf("blink messed up")
			}
			// fmt.Printf("%d: res: %v\n", i, res)
			s = res
		}
	}
}

func BenchmarkParallelBlink(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := strings.Split(input, " ")
		s := stones(Map(data, atoi))
		iterations := 25
		numWorkers := 8
		_ = parallelBlink(s, iterations, numWorkers)
	}
}
