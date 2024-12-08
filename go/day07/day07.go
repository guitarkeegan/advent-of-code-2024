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

type equation struct {
	goal    int
	numbers []int
}

func Map[T any, U any](og []T, fn func(item T) U) []U {
	res := make([]U, len(og))
	for i, val := range og {
		res[i] = fn(val)
	}
	return res
}

func load() []equation {

	var res [][]int
	atoi := func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("atoi: %v\n", err)
		}
		return num
	}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		strSlice := strings.Fields(line)
		strSlice[0] = strings.Replace(strSlice[0], ":", "", 1)
		res = append(res, Map(strSlice, atoi))
	}

	e := make([]equation, len(res))

	for i, r := range res {
		e[i] = newEquation(r)
	}

	return e
}

func newEquation(nums []int) equation {
	return equation{
		goal:    nums[0],
		numbers: nums[1:],
	}
}

type node struct {
	val   int
	left  *node
	right *node
	// part 2
	center *node
}

func newNode(val int) *node {
	return &node{
		val: val,
	}
}

func concat(n1, n2 int) int {

	s1, s2 := strconv.Itoa(n1), strconv.Itoa(n2)
	combined := fmt.Sprintf("%s%s", s1, s2)
	num, err := strconv.Atoi(combined)
	if err != nil {
		log.Fatalf("concat: %v", err)
	}

	return num
}

func compute(e equation) int {

	dbg("compute")

	n := new(node)
	n.val = e.numbers[0]
	var res int

	dbg("first node: %+v", n)

	var build func(numNode *node, idx int)
	build = func(numNode *node, idx int) {

		if numNode == nil {
			return
		}

		dbg("val: %d", numNode.val)
		dbg("goal: %d", e.goal)

		if numNode.val == e.goal && idx == len(e.numbers)-1 {
			res = e.goal
			return
		}

		if idx+1 == len(e.numbers) {
			return
		}

		lNode := newNode(numNode.val + e.numbers[idx+1])
		rNode := newNode(numNode.val * e.numbers[idx+1])
		// part 2
		cNode := newNode(concat(numNode.val, e.numbers[idx+1]))

		numNode.left = lNode
		numNode.right = rNode
		// part 2
		numNode.center = cNode

		build(numNode.left, idx+1)
		// part 2
		build(numNode.center, idx+1)
		build(numNode.right, idx+1)
	}

	dbg("call build")
	build(n, 0)

	dbg("return compute")
	return res
}

func run(e []equation) int {

	dbg("inside run")
	var res int
	for _, eq := range e {
		res += compute(eq)
	}
	dbg("end run")
	return res
}

func main() {

	dbg("load()")
	equations := load()
	dbg("run()")
	fmt.Println(run(equations))

}
