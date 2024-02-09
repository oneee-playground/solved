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
		a, b uint64

		nums [1e7 + 1]bool
	)
	nums[1] = true
	cnt := 0
	for i := uint64(2); i*i <= 1e7; i++ {
		if nums[i] {
			continue
		}
		for j := i * i; j <= 1e7; j += i {
			nums[j] = true
		}
	}
	fmt.Fscan(r, &a, &b)

	for idx := 2; idx < len(nums); idx++ {
		if nums[idx] {
			continue
		}
		num := uint64(idx)
		for j := num; j <= b/num; j *= num {
			if j*num >= a {
				cnt++
			}
		}
	}
	fmt.Fprintln(w, cnt)
}
