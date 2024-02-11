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
	var (
		n, s int
		nums []int
	)

	fmt.Fscan(r, &n, &s)
	nums = make([]int, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}

	var l int = 1e5

	right, sum := 0, nums[0]

loop:
	for left := 0; left < n; left++ {
		for right < n && sum < s {
			right++
			if right == n {
				break loop
			}
			sum += nums[right]
		}
		l = min(l, right-left+1)
		sum -= nums[left]
	}
	if l == 1e5 {
		l = 0
	}
	fmt.Fprintln(w, l)
}
