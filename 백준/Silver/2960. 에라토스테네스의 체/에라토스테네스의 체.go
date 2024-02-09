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
		nums []bool
		cnt  int
	)
	fmt.Fscan(r, &n, &k)
	nums = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if nums[i] {
			continue
		}
		cnt++
		nums[i] = true
		if cnt == k {
			fmt.Fprintln(w, i)
			return
		}
		for j := i * i; j <= n; j += i {
			if nums[j] {
				continue
			}
			nums[j] = true
			cnt++
			if cnt == k {
				fmt.Fprintln(w, j)
				return
			}
		}
	}
}
