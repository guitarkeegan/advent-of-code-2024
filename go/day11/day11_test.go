package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitEven(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		outputOne string
		outputTwo string
		outputErr error
	}{
		{"should split in two", "1234", "12", "34", nil},
		{"should return error", "12345", "", "", NonEvenErr},
		{"should eleminate zeros", "200001", "200", "1", nil},
		{"empty string should return error", "", "", "", EmptyStrErr},
	}

	str := []string{}
	s := stones(str)

	for _, test := range tests {
		s1, s2, err := s.splitEven(test.input)
		if s1 != test.outputOne || s2 != test.outputTwo || err != test.outputErr {
			t.Errorf("%s: input: %s, got: s1: %s, s2: %s, err: %s. Want: %s, %s, %s", test.name, test.input, s1, s2, err, test.outputOne, test.outputTwo, test.outputErr)
		}
	}
}

func BenchmarkBlink(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := strings.Split(input, " ")

		s := stones(data)
		for i := 0; i < 25; i++ {
			res, err := blink(s)
			if err != nil {
				fmt.Errorf("blink messed up")
			}
			s = res
		}
		fmt.Println(len(s), "stones")
	}
}
