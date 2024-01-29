package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnts := make([]int, 10)
	for n > 0 {
		cnts[n%10]++
		n /= 10
	}

	var max, replacables int
	for idx, cnt := range cnts {
		if idx == 6 || idx == 9 {
			replacables += cnt
			continue
		}
		if cnt > max {
			max = cnt
		}
	}
	if div := (replacables + 1) / 2; div > max {
		max = div
	}
	fmt.Println(max)
}
