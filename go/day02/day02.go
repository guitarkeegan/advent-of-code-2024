package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	LARGEST_LEAP = 3
)

// I will use level to maintain the state of the slice as it moves
// through the different filters
type level struct {
	value    []int
	issues   int
	ogLength int
}

// load file
func getInput(path string, levelsCh chan string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("getInput: os.Open: %s", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		levelsCh <- s.Text()
	}
	close(levelsCh)

}

// get fields
// convert to slice of int
func stringToIntSlice(strLevels chan string, intLevels chan []int) {

	for sLev := range strLevels {
		var intSlice []int
		strSlice := strings.Split(sLev, " ")
		for _, s := range strSlice {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("stringToIntSlice: strconv: %s", err)
			}
			intSlice = append(intSlice, num)
		}
		intLevels <- intSlice
	}

	close(intLevels)
}

func stringToLevel(strLevels chan string, intLevels chan *level) {

	for sLev := range strLevels {
		lev := &level{}
		strSlice := strings.Split(sLev, " ")
		for _, s := range strSlice {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("stringToIntSlice: strconv: %s", err)
			}
			lev.value = append(lev.value, num)
			lev.ogLength = len(lev.value)
		}
		fmt.Println("stringToLevel")
		fmt.Println(lev)
		intLevels <- lev
	}

	close(intLevels)
}

func isIncreasing(level []int) bool {

	// assuming always at least 2 in slice
	for i, n := range level {
		if i > 0 && n <= level[i-1] {
			return false
		}
	}
	return true
}

func isIncreasingLevel(lev *level) (*level, bool) {

	filtered := make([]int, 0, lev.ogLength)
	// assuming always at least 2 in slice
	for i, n := range lev.value {
		if i > 0 && n <= lev.value[i-1] {
			continue
		}
		filtered = append(filtered, n)
	}
	filteredLen := len(filtered)
	if filteredLen == lev.ogLength {
		return lev, true
	} else if filteredLen == lev.ogLength-1 {
		lev.issues++
		lev.value = filtered
		return lev, true
	}
	lev.issues = 0
	return nil, false
}

func isDecreasing(level []int) bool {

	// assuming always at least 2 in slice
	for i, n := range level {
		if i > 0 && n >= level[i-1] {
			return false
		}
	}
	return true
}

func isDecreasingLevel(lev *level) (*level, bool) {

	filtered := make([]int, 0, lev.ogLength)
	// assuming always at least 2 in slice
	for i, n := range lev.value {
		if i > 0 && n >= lev.value[i-1] {
			continue
		}
		filtered = append(filtered, n)
	}
	filteredLen := len(filtered)
	if filteredLen == lev.ogLength {
		return lev, true
	} else if filteredLen == lev.ogLength-1 {
		lev.issues++
		lev.value = filtered
		return lev, true
	}
	lev.issues = 0
	return nil, false
}

// check first condition or increasing
func isIncreasingOrDecreasing(intLevels chan []int, outLevels chan []int) {

	for nums := range intLevels {
		if isDecreasing(nums) || isIncreasing(nums) {
			outLevels <- nums
		}
	}
	close(outLevels)
}

// check first condition or increasing
func isIncreasingOrDecreasingLevel(inLevel chan *level, outLevel chan *level) {

	for lev := range inLevel {

		dLev, dOk := isDecreasingLevel(lev)
		if dOk {
			fmt.Println("isIncreasingOrDecreasingLevel: dLev")
			fmt.Println(dLev)
			outLevel <- dLev
			continue
		}
		iLev, iOk := isIncreasingLevel(lev)
		if iOk {
			fmt.Println("isIncreasingOrDecreasingLevel: iLev")
			fmt.Println("inc or dec")
			fmt.Println(iLev)
			outLevel <- iLev
		}
	}
	close(outLevel)
}

func absDiff(num int) int {

	if num < 0 {
		return -num
	}
	return num

}

// check second condition of step or leap
func isWithinThree(levels, levelInc chan []int) {

	for level := range levels {
		passed := true
		for i, num := range level {
			if i > 0 && (absDiff(num-level[i-1]) > LARGEST_LEAP) {
				passed = false
				break
			}
		}
		if passed {
			levelInc <- level
		}
	}
	close(levelInc)
}

func isWithinThreeLevels(inLev, outLev chan *level) {

	for lev := range inLev {

		filtered := []int{}
		for i, num := range lev.value {
			if i == 0 {
				filtered = append(filtered, num)
				continue
			}
			if i > 0 && (absDiff(num-lev.value[i-1]) <= LARGEST_LEAP) {
				filtered = append(filtered, num)
			}
		}
		filteredLen := len(filtered)
		if filteredLen >= lev.ogLength-1 {
			fmt.Println("withinThree")
			fmt.Println(lev)
			outLev <- lev
		}
	}
	close(outLev)
}

// incrementor to count passing levels
func increment(levels chan []int) int {

	var count int
	for _ = range levels {
		count++
	}
	return count
}

func incrementLevels(levels chan *level) int {

	fmt.Println("incrementLevels")

	var count int
	for l := range levels {
		fmt.Println(l)
		count++
	}
	return count

}

func part1() {
	ch1 := make(chan string)
	ch2 := make(chan []int)
	ch3 := make(chan []int)
	ch4 := make(chan []int)
	go getInput("inputs/day02", ch1)
	go stringToIntSlice(ch1, ch2)
	go isIncreasingOrDecreasing(ch2, ch3)
	go isWithinThree(ch3, ch4)
	res := increment(ch4)
	fmt.Println(res)
}

// ------ no chans!

func part2(path string) int {
	ch1 := make(chan string)
	ch2 := make(chan *level)
	ch3 := make(chan *level)
	ch4 := make(chan *level)
	go getInput(path, ch1)
	go stringToLevel(ch1, ch2)
	go isIncreasingOrDecreasingLevel(ch2, ch3)
	go isWithinThreeLevels(ch3, ch4)
	res := incrementLevels(ch4)
	return res
}

func main() {

	part1()
	fmt.Println(part2("inputs/day02"))
}
