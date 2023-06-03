package main

func solution(absolutes []int, signs []bool) int {
	sum := 0
	for i, data := range absolutes {
		if signs[i] {
			sum += data
		} else {
			sum -= data
		}
	}
	return sum
}
