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

		dp, trace []int
	)

	fmt.Fscan(r, &n)
	dp = make([]int, n+1)
	trace = make([]int, n+1)

	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		trace[i] = i - 1
		if n := dp[i/2] + 1; i%2 == 0 && dp[i] > n {
			dp[i] = n
			trace[i] = i / 2
		}
		if n := dp[i/3] + 1; i%3 == 0 && dp[i] > n {
			dp[i] = n
			trace[i] = i / 3
		}
	}

	fmt.Fprintln(w, dp[n])
	for i := n; i > 0; i = trace[i] {
		fmt.Fprintf(w, "%d ", i)
	}
}
