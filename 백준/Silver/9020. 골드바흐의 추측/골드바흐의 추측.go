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
		t, n int

		nums [1e4 + 1]bool
	)

	nums[1] = true
	for i := 2; i*i <= 1e4; i++ {
		if nums[i] {
			continue
		}
		for j := i * i; j <= 1e4; j += i {
			nums[j] = true
		}
	}

	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &n)
		first, second := 0, 0
		for i := 2; i <= n/2; i++ {
			if !nums[i] && !nums[n-i] {
				first, second = i, n-i
			}
		}
		fmt.Fprintln(w, first, second)
	}
}
