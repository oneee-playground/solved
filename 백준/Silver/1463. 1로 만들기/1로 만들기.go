package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		cnt := dp[i - 1] + 1 
		if i % 2 == 0 {
			if div2 := dp[i / 2] + 1; div2 < cnt {
				cnt = div2
			}
		}
		if i % 3 == 0 {
			if div3 := dp[i / 3] + 1; div3 < cnt {
				cnt = div3
			}
		}
		dp[i] = cnt
	}
	fmt.Println(dp[n])
}