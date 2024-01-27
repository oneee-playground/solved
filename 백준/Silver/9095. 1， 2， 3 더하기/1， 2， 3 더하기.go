package main

import (
	"fmt"
)

func main() {
	dp := [12]int{1, 1, 2, 4}
	for i := 4; i < 12; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		fmt.Println(dp[n])
	}
}
