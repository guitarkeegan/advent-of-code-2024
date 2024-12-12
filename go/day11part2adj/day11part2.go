package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
type stoneMap map[int]stones

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
	if n == 0 {
		return -1, fmt.Errorf("replaceStone: can't be zero: n: %d", n)
	}
	numDigits := int(math.Log10(float64(n))) + 1
	if numDigits%2 == 0 {
		return -1, fmt.Errorf("replaceStone: n cannot be even length, length: %d, for n: %d", numDigits, n)
	}
	return n * TwentyTwentyFour, nil

}

func bfs(s stones, sm *stoneMap, iterations int) (int, error) {

	dbg("bfs")

	if len(s) == 0 {
		return 0, nil
	}

	q := stones([]int{})
	q = append(q, s...)

	count := len(s)

	for i := 0; i < iterations; i++ {
		dbg("outer for")

		length := len(q)
		for length > 0 {
			stn := q[:1]
			q = q[1:]
			stns, err := getOrCreateChildren(stn, sm)
			if err != nil {
				return -1, fmt.Errorf("bfs: getOrCreateChildren: %w", err)
			}
			q = append(q, stns...)
			length--
		}
		count = len(q)
	}

	return count, nil
}

func getOrCreateChildren(s stones, sm *stoneMap) (stones, error) {
	dbg("getOrCreateChildren")

	res := stones{}
	for _, stone := range s {

		if stn, ok := (*sm)[stone]; ok {
			dbg("already in map")
			res = append(res, stn...)
			continue
		}

		if stone == 0 {
			one, err := s.zeroToOne(stone)
			if err != nil {
				return stones([]int{}), fmt.Errorf("getOrCreateChildren: zeroToOne: %w", err)
			}
			(*sm)[stone] = stones([]int{one})
			res = append(res, one)
			dbg("  append res: %v", res)
			continue
		}

		numLength := int(math.Log10(float64(stone)) + 1)
		if numLength%2 == 0 {
			s1, s2, err := s.splitEven(stone)
			if err != nil {
				return stones([]int{}), fmt.Errorf("getOrCreateChildren: splitEven: %w", err)
			}
			(*sm)[stone] = stones([]int{s1, s2})
			res = append(res, s1, s2)
		} else {
			newStone, err := s.replaceStone(stone)
			if err != nil {
				return stones([]int{}), fmt.Errorf("getOrCreateChildren: replaceStone: %w", err)
			}
			(*sm)[stone] = stones([]int{newStone})
			res = append(res, newStone)
		}
		dbg("res: %v", res)

	}
	return res, nil
}

func blink(s stones, iterations int) (int, error) {
	dbg("blink")
	if len(s) == 0 {
		return -1, fmt.Errorf("blink: you got no stones!")
	}

	// place the first stone's children in the map
	sm := make(stoneMap)
	_, err := getOrCreateChildren(s, &sm)
	if err != nil {
		return -1, fmt.Errorf("blink: getOrCreateChildren: %w", err)
	}

	length, err := bfs(s, &sm, iterations)
	if err != nil {
		return -1, fmt.Errorf("blink: bfs: %w", err)
	}
	dbg("stoneMap: %+v", sm)
	return length, nil

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

	data := strings.Split(input, " ")
	s := stones(Map(data, atoi))
	res, err := blink(s, 75)
	if err != nil {
		fmt.Errorf("problem with blink: %v", err)
	}
	fmt.Println(res, "stones")

}
