package golab

import (
	"math"
	"strconv"
)

var (
	min = 1<<31 - 1
	max = -1 << 31
)

func solution(s string) string {

	buf := make([]rune, 0, 100)

	for _, r := range s {
		if r != ' ' {
			buf = append(buf, r)
			continue
		}

		calculate(buf)
		buf = nil
	}

	calculate(buf)

	return strconv.Itoa(min) + " " + strconv.Itoa(max)
}

func calculate(buf []rune) {

	isNegative := buf[0] == '-'

	if isNegative {
		buf = buf[1:]
	}

	var sum int
	for idx, r := range buf {
		sum += int(r-'0') * int(math.Pow10(len(buf)-idx-1))
	}

	if isNegative {
		sum = -sum
	}

	if sum > max {
		max = sum
	}

	if sum < min {
		min = sum
	}
}
