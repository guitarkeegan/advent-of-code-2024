package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map[T any, U any](input []T, fn func(T) U) []U {
	res := make([]U, len(input))
	for _, val := range input {
		res = append(res, fn(val))
	}
	return res
}

func load(input string) [][]int {
	rows := strings.Split(strings.TrimSpace(input), "\n")

	// Prepare a 2D slice to hold the rows and their integer elements
	var matrix [][]int

	// Process each row
	for _, row := range rows {
		// Split the row into string elements
		stringElements := strings.Fields(row)

		// Convert string elements to integers
		var intElements []int
		for _, str := range stringElements {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting to int: %v\n", err)
				continue
			}
			intElements = append(intElements, num)
		}

		// Append the integer row to the matrix
		matrix = append(matrix, intElements)
	}

	return matrix

}

type leveltron struct {
	s1    []int
	s2    []int
	tally map[string]int
	min   int
	max   int
}

func NewLeveltron(stack1, stack2 []int) *leveltron {
	return &leveltron{
		s1: stack1,
		s2: stack2,
		tally: map[string]int{
			"gt": 0,
			"lt": 0,
		},
		min: 1,
		max: 3,
	}
}

func (lt *leveltron) insert(n int) (bool, error) {

	if len(lt.s1) == 0 {
		lt.s1 = append(lt.s1, n)
		return true, nil
	}

	diff := lt.getAbsDiff(lt.s1[len(lt.s1)-1], n)
	if diff > lt.min && diff < lt.max {
		lt.s1 = append(lt.s1, n)
	}
}

func (lt *leveltron) getAbsDiff(n1, n2 int) int {
	diff := n1 - n2
	if diff < 0 {
		return -diff
	}
	return diff
}

func main() {

	rows := load(testInput)

}
