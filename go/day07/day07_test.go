package main

import "testing"

func TestCompute(t *testing.T) {

	tests := []struct {
		name     string
		input    equation
		expected int
	}{
		{"3 + 5", equation{
			goal:    8,
			numbers: []int{3, 5},
		}, 8},
		{"3 + 5 * 2", equation{
			goal:    16,
			numbers: []int{3, 5, 2},
		}, 16},
		{"3 * 5 + 2 || 1", equation{
			goal:    171,
			numbers: []int{3, 5, 2, 1},
		}, 171},
		{"3 || 5 + 2 || 1", equation{
			goal:    371,
			numbers: []int{3, 5, 2, 1},
		}, 371},
		{"3 * 5 || 2 + 1", equation{
			goal:    153,
			numbers: []int{3, 5, 2, 1},
		}, 153},
		{"shouldn't work", equation{
			goal:    154,
			numbers: []int{3, 5, 2, 1},
		}, 0},
		{"no way to get this", equation{
			goal:    19,
			numbers: []int{4, 4, 4},
		}, 0},
		{"shouldn't work double digits", equation{
			goal:    62,
			numbers: []int{10, 10, 10},
		}, 0},
		{"shouldn't work triple digits", equation{
			goal:    21,
			numbers: []int{100, 100, 100},
		}, 0},
		{"might finish early", equation{
			goal:    3,
			numbers: []int{1, 2, 3},
		}, 0},
	}

	for _, test := range tests {
		res := compute(test.input)
		if test.expected != res {
			t.Errorf("Got: %d, Want: %d", res, test.expected)
		}
	}

}
