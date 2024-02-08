package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, k  int
		coins []int
		cnt   int
	)

	fmt.Fscan(r, &n, &k)
	coins = make([]int, n)
	for i := range coins {
		fmt.Fscan(r, &coins[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	for _, coin := range coins {
		for coin <= k {
			k -= coin
			cnt++
		}
		if k == 0 {
			break
		}
	}
	fmt.Fprintln(w, cnt)
}
