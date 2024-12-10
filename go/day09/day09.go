package main

import (
	"fmt"
	"log"
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

const EMPTY = -1

func load() []int {

	diskMapStr := strings.Split(testInput, "")
	// try pointers to get nil at the empty slots
	diskMap := make([]int, len(diskMapStr))
	for i, ch := range diskMapStr {
		num, err := strconv.Atoi(ch)
		if err != nil {
			log.Fatalf("atoi: %v", err)
		}
		diskMap[i] = num
	}
	return diskMap

}

func getMaxLength(diskMap []int) int {
	var res int
	for _, num := range diskMap {
		res += num
	}

	return res
}

func main() {

	diskMap := load()
	mapLength := getMaxLength(diskMap)
	denseFormat := make([]int, mapLength)
	dbg("mapLength: %d", mapLength)
	var id int
	var p1 int

	for i, num := range diskMap {
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				denseFormat[p1] = id
				p1++
			}
		} else {
			for j := 0; j < num; j++ {
				denseFormat[p1] = EMPTY
				p1++
			}
			id++
		}
	}

	// 2333133121414131402
	// 00...111...2...333.44.5555.6666.777.888899
	diskMapLeft, diskMapRight := 0, len(diskMap)-1
	denseLeft, denseRight := 0, len(denseFormat)
	for i := 0; i < 1; i++ {
		dbg("denseFormat: %v", denseFormat)
		dbg("diskMap: %v", diskMap)
		dbg("diskMapLeft: %d, diskMapRight: %d", diskMapLeft, diskMapRight)
		dbg("denseLeft: %d, denseRight: %d", denseLeft, denseRight)
		if diskMapLeft%2 == 0 {
			denseLeft += diskMap[diskMapLeft]
			diskMapLeft++
		}
	}

	// dbg("before condense: denseFormat: %v", denseFormat)
	// condense
	// part 1

	//	var part1 = func() {
	//		p1 = 0
	//		p2 := mapLength - 1
	//		for p1 < p2 {
	//			if denseFormat[p1] != EMPTY {
	//				p1++
	//				continue
	//			}
	//			if denseFormat[p2] == EMPTY {
	//				p2--
	//				continue
	//			}
	//			denseFormat[p1] = denseFormat[p2]
	//			denseFormat[p2] = EMPTY
	//			p1++
	//			p2--
	//		}
	//	}
	//
	//	part1()

	// assume that it ends with a file?
	// dbg("after condense: denseFormat: %v", denseFormat)

	// getChecksum
	var checkSum int
	for i, n := range denseFormat {
		if n == EMPTY {
			// part 1
			// break
			// part 2
			continue
		}
		checkSum += i * n
	}
	fmt.Println(checkSum)

}
