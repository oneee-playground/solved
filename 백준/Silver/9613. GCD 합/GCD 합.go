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
		t, n int
		nums []int
	)
	fmt.Fscan(r, &t)
	for tc := 0; tc < t; tc++ {
		fmt.Fscan(r, &n)
		nums = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &nums[i])
		}

		sum := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				sum += gcd(nums[i], nums[j])
			}
		}
		fmt.Fprintln(w, sum)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
