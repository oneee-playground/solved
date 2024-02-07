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

		stairs []int
		dp     []int

		sum int
	)

	fmt.Fscan(r, &n)
	stairs = make([]int, n)
	dp = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &stairs[i])
		sum += stairs[i]
	}

	if n <= 2 {
		fmt.Fprintln(w, sum)
		return
	}
	dp[0] = stairs[0]
	dp[1] = stairs[1]
	dp[2] = stairs[2]
	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-2], dp[i-3]) + stairs[i]
	}
	fmt.Fprintln(w, sum-min(dp[n-2], dp[n-3]))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
