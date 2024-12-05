package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var dbg = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(string, ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format+"\n", as...)
	}
}()

// load the lines
func loadLines(path string, lines chan string) {

	dbg("loadLines")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		lines <- s.Text()
	}
	close(lines)
}

// takes in input string chan, and and output []string chan
// extracts the regexp matches
// sends them to extract numbers
func extractMatches(inCh chan string, outCh chan []string) {

	dbg("extractMatches")
	re, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		log.Fatal(err)
	}

	for line := range inCh {
		outCh <- re.FindAllString(line, -1)
	}
	close(outCh)
}

func filterMuls(strArr []string) []string {

	var res []string
	on := true
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)
	mulPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for _, char := range strArr {

		if doPattern.MatchString(char) {
			on = true
		}

		if dontPattern.MatchString(char) {
			dbg("don't pattern")
			on = false
		}

		if on {
			if mulPattern.MatchString(char) {
				res = append(res, char)
			}
		}
	}

	return res

}

// extractornator
// iterates through the input
// turns sends pairs to outCh only when do() is enables
// takes string chan and []string chan as params
func extractornator(inCh chan string, outCh chan []string) {

	var (
		pattern = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	)

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	for s := range inCh {
		cmds := re.FindAllString(s, -1)
		dbg("cmds: %v\n", cmds)
		outCh <- filterMuls(cmds)
	}
	close(outCh)

}

// util takes []string and returns a [][]int of pairs
func getNumUtil(strArr []string) [][]int {

	pairs := make([][]int, len(strArr))
	re, err := regexp.Compile(`\d{1,3}`)
	if err != nil {
		log.Fatal(err)
	}

	for i, str := range strArr {
		numStrs := re.FindAllString(str, -1)
		var pair []int
		for _, s := range numStrs {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			pair = append(pair, n)
		}
		pairs[i] = pair
	}
	return pairs
}

// extract the numbers
// takes inputs []string chan, and []int chan params
// places the number pair into a length 2 array
// and sends to the []int chan
func extractNumbers(strCh chan []string, intCh chan [][]int) {

	dbg("extractNumbers")
	for str := range strCh {
		intCh <- getNumUtil(str)
	}
	close(intCh)

}

func mulPairUtil(pair []int) int {
	return pair[0] * pair[1]
}

// multiply the pair and send to addCh
// takes [][]int chan, and int chan params
func mulPairs(pairs chan [][]int, addCh chan int) {

	dbg("mulPairs")
	for pair := range pairs {
		for _, p := range pair {
			// dbg("mulPairs: %v", p)
			addCh <- mulPairUtil(p)
		}
	}
	close(addCh)
}

// sum the results and return result
func sumResults(nCh chan int) int {
	var total int

	dbg("sumResults")
	for n := range nCh {
		total += n
	}

	return total
}

func part1(path string) int {

	ch1 := make(chan string)
	ch2 := make(chan []string)
	ch3 := make(chan [][]int)
	ch4 := make(chan int)
	go loadLines(path, ch1)
	go extractMatches(ch1, ch2)
	go extractNumbers(ch2, ch3)
	go mulPairs(ch3, ch4)
	return sumResults(ch4)
}

func part2(path string) int {

	ch1 := make(chan string)
	ch2 := make(chan []string)
	ch3 := make(chan [][]int)
	ch4 := make(chan int)
	go loadLines(path, ch1)
	go extractornator(ch1, ch2) // alt module
	go extractNumbers(ch2, ch3)
	go mulPairs(ch3, ch4)
	return sumResults(ch4)
}

func main() {
	// fmt.Println(part1("inputs/day03"))
	fmt.Println(part2("inputs/day03"))
}
