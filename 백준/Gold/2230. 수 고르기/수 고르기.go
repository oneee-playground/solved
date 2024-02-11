package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m int
		nums []int
	)

	fmt.Fscan(r, &n, &m)
	nums = make([]int, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}

	sort.Ints(nums)

	var minDiff int = 2e9
	right := 1
	for left := 0; left < n; left++ {
		for right < n {
			if abs(nums[left]-nums[right]) >= m {
				minDiff = min(minDiff, abs(nums[left]-nums[right]))
				break
			}
			right++
		}
	}
	fmt.Fprintln(w, minDiff)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
