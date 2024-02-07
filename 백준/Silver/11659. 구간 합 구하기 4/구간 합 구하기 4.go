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
		dp   []int
	)

	fmt.Fscan(r, &n, &m)
	dp = make([]int, n+1)
	fmt.Fscan(r, &dp[1])
	for i := 2; i <= n; i++ {
		fmt.Fscan(r, &dp[i])
		dp[i] += dp[i-1]
	}

	for i := 0; i < m; i++ {
		var start, end int
		fmt.Fscan(r, &start, &end)
		fmt.Fprintln(w, dp[end]-dp[start-1])
	}
}
