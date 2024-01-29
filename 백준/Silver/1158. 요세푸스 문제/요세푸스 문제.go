package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]bool, n)
	res := make([]int, 0, n)
	idx, cnt, end := -1, 0, 0
	for end < n {
		cnt++
		for {
			idx = (idx + 1) % n
			if !arr[idx] {
				break
			}
		}
		if cnt == k {
			arr[idx] = true
			res = append(res, idx+1)
			cnt = 0
			end++
		}
	}
	fmt.Print("<")
	for idx, n := range res {
		fmt.Print(n)
		if idx != len(res)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(">")
}
