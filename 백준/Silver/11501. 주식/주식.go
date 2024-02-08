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
		tc, n  int
		prices []int
	)

	fmt.Fscan(r, &tc)
	for t := 0; t < tc; t++ {
		fmt.Fscan(r, &n)
		for i := 0; i < n; i++ {
			var p int
			fmt.Fscan(r, &p)
			prices = append(prices, p)
		}

		total, max := 0, 0
		for i := n - 1; i >= 0; i-- {
			v := prices[i]
			if v > max {
				max = v
			}
			if v < max {
				total += max - v
			}
		}

		fmt.Fprintln(w, total)
		prices = prices[:0]
	}
}
