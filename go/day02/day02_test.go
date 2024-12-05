package main

import (
	"os"
	"reflect"
	"testing"
)

func TestIsIncreasing(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		expect bool
	}{
		{"is increasing", []int{1, 2, 3, 4}, true},
		{"is increasing by 2", []int{1, 3, 5, 7}, true},
		{"has repeated number", []int{1, 2, 2, 4}, false},
		{"changes direction", []int{1, 2, 1, 4}, false},
		{"is decreasing", []int{3, 2, 1}, false},
	}

	for _, test := range tests {
		res := isIncreasing(test.input)
		if test.expect != res {
			t.Errorf("%s: Got: %t, Want: %t, slice: %v\n", test.name, res, test.expect, test.input)
		}
	}
}

func TestIsWithinThree(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"increase by 1", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"increase by 2", []int{1, 3, 5, 7}, []int{1, 3, 5, 7}},
		{"increase by 5 (filtered out)", []int{0, 5, 10, 15}, nil}, // Filtered case
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create input and output channels
			ch1 := make(chan []int, 1) // Buffered to allow sending without blocking
			ch2 := make(chan []int)    // Unbuffered for output

			go isWithinThree(ch1, ch2) // Start function in a goroutine

			// Send input and close the channel
			ch1 <- test.input
			close(ch1)

			// Collect output
			var output []int
			select {
			case result, ok := <-ch2:
				if ok {
					output = result
				} else {
					output = nil // Channel closed without sending a value
				}
			}

			// Compare output
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("GOT: %v, WANT: %v", output, test.expected)
			}
		})
	}
}

func TestStringToLevel(t *testing.T) {

	lev := &level{
		value:    []int{48, 51, 52, 53, 52},
		issues:   0,
		ogLength: 5,
	}
	test := struct {
		name   string
		str    string
		expect *level
	}{"should ouput a new level", "48 51 52 53 52", lev}

	ch1 := make(chan string)
	ch2 := make(chan *level)

	go stringToLevel(ch1, ch2)

	ch1 <- test.str
	close(ch1)

	select {
	case l := <-ch2:
		if !reflect.DeepEqual(l.value, test.expect.value) {
			t.Errorf("Want: %v, Got: %v", test.expect.value, l.value)
		}
		if l.ogLength != test.expect.ogLength {
			t.Errorf("Want: %d, Got: %d", test.expect.ogLength, l.ogLength)
		}
	}

}

func TestPart2(t *testing.T) {

	err := os.Chdir("../.")
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	test := struct {
		name      string
		inputPath string
		output    int
	}{"2 should pass", "inputs/day02-test", 2}

	res := part2(test.inputPath)
	if res != test.output {
		t.Errorf("%s: Got: %d, Want: %d\n", test.name, res, test.output)
	}
}
