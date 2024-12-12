package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

var dbg = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(string, ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format+"\n", as...)
	}
}()

var (
	NonEvenErr       = errors.New("input must be even number")
	EmptyStrErr      = errors.New("string cannot be empty")
	TwentyTwentyFour = 2024
)

var testInput = strings.TrimSpace(`
125 17
	`)
var input = strings.TrimSpace(`0 7 198844 5687836 58 2478 25475 894`)

type stones []int

func (s stones) zeroToOne(zero int) (int, error) {
	if zero == 0 {
		return 1, nil
	}
	return -1, fmt.Errorf("can only pass in '0'. passed in: %d", zero)
}

// AI helper
func (s stones) splitEven(num int) (int, int, error) {
	numDigits := int(math.Log10(float64(num))) + 1
	if numDigits%2 != 0 {
		return -1, -1, NonEvenErr
	}

	// Calculate the divisor
	divisor := int(math.Pow(10, float64(numDigits/2)))

	// Split the number
	left := num / divisor
	right := num % divisor

	return left, right, nil
}

func (s stones) replaceStone(n int) (int, error) {
	dbg("replaceStone")
	dbg("  n: %d", n)
	if n == 0 {
		return -1, fmt.Errorf("replaceStone: can't be zero: n: %d", n)
	}
	numDigits := int(math.Log10(float64(n))) + 1
	dbg("  numDigits: %d", numDigits)
	if numDigits%2 == 0 {
		return -1, fmt.Errorf("replaceStone: n cannot be even length, length: %d, for n: %d", numDigits, n)
	}
	dbg("  returning: n * 2024: %d * %d", n, TwentyTwentyFour)
	return n * TwentyTwentyFour, nil

}

func blink(s stones) (stones, error) {
	if len(s) == 0 {
		return stones([]int{}), fmt.Errorf("blink: you got no stones!")
	}
	var res stones
	dbg("blink: s: %v", s)
	dbg("blink: res: %v", res)
	for _, stone := range s {
		if stone == 0 {
			dbg("  stone is zero")
			one, err := s.zeroToOne(stone)
			if err != nil {
				return stones([]int{}), fmt.Errorf("blink: zeroToOne: %w", err)
			}
			res = append(res, one)
			dbg("  append res: %v", res)
			continue
		}

		numLength := int(math.Log10(float64(stone)) + 1)
		dbg("numLength for stone: %d, is %d", stone, numLength)
		if numLength%2 == 0 {
			dbg("  stone is even length")
			s1, s2, err := s.splitEven(stone)
			if err != nil {
				return stones([]int{}), fmt.Errorf("blink: splitEven: %w", err)
			}
			res = append(res, s1, s2)
		} else {
			dbg("  replace stone")
			newStone, err := s.replaceStone(stone)
			dbg("  newStone: %d", newStone)
			if err != nil {
				return stones([]int{}), fmt.Errorf("blink: replaceStone: %w", err)
			}
			res = append(res, newStone)
			dbg("    newStone: %d, res: %v", newStone, res)
		}

	}
	return res, nil
}

func parallelBlink(input stones, iterations, numWorkers int) []int {
	current := input
	for it := 0; it < iterations; it++ {
		chunkSize := (len(current) + numWorkers - 1) / numWorkers
		results := make([]stones, numWorkers)
		var wg sync.WaitGroup
		wg.Add(numWorkers)

		for i := 0; i < numWorkers; i++ {
			start := i * chunkSize
			end := start + chunkSize
			if end > len(current) {
				end = len(current)
			}

			// Avoid spawning goroutines with empty slices
			if start >= len(current) {
				wg.Done()
				continue
			}

			go func(idx int, chunk stones) {
				defer wg.Done()
				res, err := blink(chunk)
				if err != nil {
					fmt.Errorf("blink failed in go routine\n")
					return
				}
				results[idx] = res
			}(i, current[start:end])
		}

		wg.Wait()

		next := make(stones, 0, len(current)*2)
		for _, result := range results {
			next = append(next, result...)
		}
		current = next
	}
	return current
}

func Map[U, T any](og []U, fn func(u U) T) []T {
	res := make([]T, len(og))
	for i, item := range og {
		res[i] = fn(item)
	}
	return res
}

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("atoi: failed to convert: %s, reason: %v", s, err)
	}
	return num
}

func main() {

	data := strings.Split(testInput, " ")
	s := stones(Map(data, atoi))
	iterations := 75
	numWorkers := 8
	res := parallelBlink(s, iterations, numWorkers)
	fmt.Println(len(res), "stones")
}
