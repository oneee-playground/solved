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
		n int

		red, green, blue []int

		dp [][3]int
	)

	fmt.Fscan(r, &n)
	red = make([]int, n)
	green = make([]int, n)
	blue = make([]int, n)
	dp = make([][3]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &red[i], &green[i], &blue[i])
	}

	dp[0][0] = red[0]
	dp[0][1] = green[0]
	dp[0][2] = blue[0]
	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + red[i]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + green[i]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + blue[i]
	}
	fmt.Fprintln(w, min(dp[n-1][0], dp[n-1][1], dp[n-1][2]))
}

func min(nums ...int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m
}
