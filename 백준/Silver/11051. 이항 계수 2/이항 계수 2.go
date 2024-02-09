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
		n, k int
		dp   [][]int
	)

	fmt.Fscan(r, &n, &k)
	dp = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0], dp[i][i] = 1, 1
		for j := 1; j < i; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]) % 10007
		}
	}
	fmt.Fprintln(w, dp[n][k])
}
