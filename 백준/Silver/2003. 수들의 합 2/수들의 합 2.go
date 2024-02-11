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
		n, m int
		nums []int
	)

	fmt.Fscan(r, &n, &m)
	nums = make([]int, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}
	cnt := 0
	right, sum := 0, nums[0]
loop:
	for left := 0; left < len(nums); left++ {
		for right < len(nums) && sum < m {
			right++
			if right == len(nums) {
				break loop
			}
			sum += nums[right]
		}
		if sum == m {
			cnt++
		}
		sum -= nums[left]
	}
	fmt.Fprintln(w, cnt)
}
