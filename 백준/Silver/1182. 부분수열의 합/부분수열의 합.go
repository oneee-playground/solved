package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &s)
	nums = make([]int, n)
	check = make([]bool, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}
	find(0)
	if s == 0 {
		cnt--
	}
	fmt.Fprintln(w, cnt)
}

var (
	n, s, cnt int
	nums      []int
	check     []bool
)

func find(i int) {
	if i == n {
		sum := 0
		for idx, num := range nums {
			if check[idx] {
				sum += num
			}
		}
		if sum == s {
			cnt++
		}
		return
	}

	check[i] = true
	find(i + 1)
	check[i] = false
	find(i + 1)
}
