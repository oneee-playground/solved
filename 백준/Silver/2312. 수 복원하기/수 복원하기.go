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

		nums [1e7 + 1]bool
	)

	nums[1] = true
	for i := 2; i*i <= 1e7; i++ {
		if nums[i] {
			continue
		}
		for j := i * i; j <= 1e7; j += i {
			nums[j] = true
		}
	}

	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &n)
		for i := 2; i*i <= n; i++ {
			cnt := 0
			for n%i == 0 {
				cnt++
				n /= i
			}
			if cnt > 0 {
				fmt.Fprintln(w, i, cnt)
			}
		}
		if n != 1 {
			fmt.Fprintln(w, n, 1)
		}
	}
}
