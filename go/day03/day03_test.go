package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestExtractMatches(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		expected []string
	}{
		{"should return 2 ops", "kjamul(12,543)j>k3!mu8mul(4,54)", []string{"mul(12,543)", "mul(4,54)"}},
		{"should return 1 op", "mul83kmulas3mul(98kas?mul(123,123)", []string{"mul(123,123)"}},
		{"should return 0 ops", "mul83kmulas3mul(98kas?mul(18kaj,123)", nil},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1) // Increment wait group counter
		ch1 := make(chan string)
		ch2 := make(chan []string)

		go func(input string, expected []string, name string) {
			defer wg.Done() // Decrement wait group counter when done
			go extractMatches(ch1, ch2)
			ch1 <- input
			close(ch1) // Close the input channel to signal completion

			select {
			case res := <-ch2:
				if !reflect.DeepEqual(res, expected) {
					t.Errorf("Test: %s, Got: %v, Want: %v\n", name, res, expected)
				}
			}
		}(test.inputStr, test.expected, test.name)
	}

	wg.Wait() // Wait for all goroutines to complete
}

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []string
		expected [][]int
	}{
		{"one item", []string{"mul(1,1)"}, [][]int{{1, 1}}},
		{"two items", []string{"mul(1,1)", "mul(1,2)"}, [][]int{{1, 1}, {1, 2}}},
		{"three items", []string{"mul(1,1)", "mul(1,2)", "mul(10,2)"}, [][]int{{1, 1}, {1, 2}, {10, 2}}},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1) // Increment wait group counter
		ch1 := make(chan []string)
		ch2 := make(chan [][]int)

		go func(input []string, expected [][]int, name string) {
			defer wg.Done() // Decrement wait group counter when done
			go extractNumbers(ch1, ch2)
			ch1 <- input
			close(ch1) // Close the input channel to signal completion

			select {
			case res := <-ch2:
				if !reflect.DeepEqual(res, expected) {
					t.Errorf("Test: %s, Got: %v, Want: %v\n", name, res, expected)
				}
			}
		}(test.inputArr, test.expected, test.name)
	}

	wg.Wait() // Wait for all goroutines to complete
}

func TestFilterMuls(t *testing.T) {

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"should output 1", []string{"mul(1,1)"}, []string{"mul(1,1)"}},
		{"should output 1", []string{"mul(1,1)", "don't()", "mul(1,1)"}, []string{"mul(1,1)"}},
		{"should output 2", []string{"mul(1,1)", "don't()", "mul(1,1)", "do()", "mul(1,1)"}, []string{"mul(1,1)", "mul(1,1)"}},
	}

	for _, test := range tests {
		res := filterMuls(test.input)
		if !reflect.DeepEqual(test.expected, res) {
			t.Fatalf("TestName: %s: Got: %v, Want: %v", test.name, res, test.expected)
		}
	}
}

func TestExtractonator(t *testing.T) {

	tests := []struct {
		name     string
		inputStr string
		expected []string
	}{
		{"should return 2 ops", "kjamul(12,543)j>k3!mu8mul(4,54)", []string{"mul(12,543)", "mul(4,54)"}},
		{"should return 1 op", "mul83kmulas3mul(98kas?mul(123,123)", []string{"mul(123,123)"}},
		{"should return 0 ops", "mul83kmulas3mul(98kas?mul(18kaj,123)", nil},
		{"should give the ", "don't()[:from()@{who() ?mul(305,335)[when()when();where(751,621)what()mul(395,86)how()?,who():>mul(349,362)how()?", nil},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1) // Increment wait group counter
		ch1 := make(chan string)
		ch2 := make(chan []string)

		go func(input string, expected []string, name string) {
			defer wg.Done() // Decrement wait group counter when done
			go extractornator(ch1, ch2)
			ch1 <- input
			close(ch1) // Close the input channel to signal completion

			select {
			case res := <-ch2:
				if !reflect.DeepEqual(res, expected) {
					t.Errorf("Test: %s, Got: %v, Want: %v\n", name, res, expected)
				}
			}
		}(test.inputStr, test.expected, test.name)
	}

	wg.Wait() // Wait for all goroutines to complete
}
