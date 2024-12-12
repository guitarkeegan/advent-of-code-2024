package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	NonEvenErr       = errors.New("input must be even number")
	EmptyStrErr      = errors.New("string cannot be empty")
	TwentyTwentyFour = 2024
)

var testInput = strings.TrimSpace(`
125 17
	`)
var input = strings.TrimSpace(`0 7 198844 5687836 58 2478 25475 894`)

type stones []string

func (s stones) zeroToOne(zero string) (string, error) {
	if zero == "0" {
		return "1", nil
	}
	return "", fmt.Errorf("can only pass in '0'. passed in: %s", zero)
}

func (s stones) splitEven(evenNum string) (string, string, error) {
	if evenNum == "" {
		return "", "", EmptyStrErr
	}
	if len(evenNum)%2 == 0 {
		left, right := evenNum[:len(evenNum)/2], evenNum[len(evenNum)/2:]
		nLeft, err := strconv.Atoi(left)
		if err != nil {
			return "", "", fmt.Errorf("splitEven nLeft: Atoi: %v", err)
		}
		nRight, err := strconv.Atoi(right)
		if err != nil {
			return "", "", fmt.Errorf("splitEven nRight: Atoi %v", err)
		}
		// back to string
		return strconv.Itoa(nLeft), strconv.Itoa(nRight), nil
	}
	return "", "", NonEvenErr
}

func (s stones) replaceStone(n string) (string, error) {
	if n == "0" || len(n)%2 == 0 {
		return "", fmt.Errorf("stone is 0 or even: %s", n)
	}
	num, err := strconv.Atoi(n)
	if err != nil {
		return "", fmt.Errorf("replaceStone: Atoi: %v", err)
	}
	newNum := num * TwentyTwentyFour
	return strconv.Itoa(newNum), nil
}

func (s stones) len() int {
	return len(s)
}

func blink(s stones) (stones, error) {
	if len(s) == 0 {
		return stones([]string{}), fmt.Errorf("blink: you got no stones!")
	}
	var res stones
	for _, stone := range s {
		if stone == "0" {
			one, err := s.zeroToOne(stone)
			if err != nil {
				return stones([]string{}), fmt.Errorf("blink: zeroToOne: %w", err)
			}
			res = append(res, one)
		} else if len(stone)%2 == 0 {
			s1, s2, err := s.splitEven(stone)
			if err != nil {
				return stones([]string{}), fmt.Errorf("blink: splitEven: %w", err)
			}
			res = append(res, s1, s2)
		} else {
			newStone, err := s.replaceStone(stone)
			if err != nil {
				return stones([]string{}), fmt.Errorf("blink: replaceStone: %w", err)
			}
			res = append(res, newStone)
		}

	}
	return res, nil
}

func main() {

	data := strings.Split(input, " ")

	s := stones(data)
	for i := 0; i < 75; i++ {
		res, err := blink(s)
		if err != nil {
			fmt.Errorf("blink messed up")
		}
		s = res
	}
	fmt.Println(len(s), "stones")

}
