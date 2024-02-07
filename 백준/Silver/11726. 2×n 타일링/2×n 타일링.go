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
		n  int
		dp []int
	)

	fmt.Fscan(r, &n)
	dp = make([]int, n)
	if n <= 2 {
		fmt.Fprintln(w, n)
		return
	}
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10007
	}
	fmt.Fprintln(w, dp[n-1])
}
