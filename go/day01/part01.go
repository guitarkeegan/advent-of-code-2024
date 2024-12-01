package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func getLeftRightValues(path string) ([]int, []int) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open: %s", err)
	}

	var leftCol []int
	var rightCol []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		left, err := strconv.Atoi(scanner.Text()[:5])
		if err != nil {
			log.Fatalf("scannerScan: ParseFloat: left: %s", err)
		}
		right, err := strconv.Atoi(scanner.Text()[8:13])
		if err != nil {
			log.Fatalf("scannerScan: ParseFloat: right: %s", err)
		}
		leftCol = append(leftCol, left)
		rightCol = append(rightCol, right)
	}

	return leftCol, rightCol

}

func getAbs(n int) int {

	if n < 0 {
		return -n
	}
	return n
}

func getDiffs(leftCol []int, rightCol []int) int {

	var res int

	for i := range leftCol {

		res += getAbs(leftCol[i] - rightCol[i])
	}

	return res
}

func part1() {
	// Part I --------------
	// make two arrays
	leftCol, rightCol := getLeftRightValues("inputs/day01")
	// sort arrays
	sort.Ints(leftCol)
	sort.Ints(rightCol)
	// get absolute diff for each, from smallest to biggest
	res := getDiffs(leftCol, rightCol)
	fmt.Println(res)
}

func getInput(path string) ([]int, map[int]int) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("mapify: Open: %s", err)
	}

	var leftCol []int
	rightMap := make(map[int]int)

	s := bufio.NewScanner(file)

	for s.Scan() {
		left, err := strconv.Atoi(s.Text()[:5])
		if err != nil {
			log.Fatalf("sScan: ParseFloat: left: %s", err)
		}
		right, err := strconv.Atoi(s.Text()[8:13])
		if err != nil {
			log.Fatalf("scannerScan: ParseFloat: right: %s", err)
		}
		leftCol = append(leftCol, left)
		rightMap[right] += 1
	}

	return leftCol, rightMap
}

func getSimularityScore(leftCol []int, rMap map[int]int) int {

	var result int

	for _, num := range leftCol {
		result += rMap[num] * num
	}

	return result
}

func part2() {
	lc, rm := getInput("inputs/day01")
	res := getSimularityScore(lc, rm)
	fmt.Println(res)
}

func main() {

	// part1()
	// Part II -----------------
	part2()
}
