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
		t int

		nums [1e8 + 1]bool
		tc   []int
		max  int
	)

	fmt.Fscan(r, &t)
	tc = make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &tc[i])
		if tc[i] > max {
			max = tc[i]
		}
	}

	nums[1] = true
	for i := 2; i*i <= max; i++ {
		if nums[i] {
			continue
		}
		for j := i * i; j <= max; j += i {
			nums[j] = true
		}
	}

	for _, n := range tc {
		cnt := 0
		for i := 2; i <= n/2; i++ {
			if !nums[i] && !nums[n-i] {
				cnt++
			}
		}
		fmt.Fprintln(w, cnt)
	}
}
