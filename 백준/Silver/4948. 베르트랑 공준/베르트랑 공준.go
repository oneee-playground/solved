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
		n    int
		nums [(2 * 123_456) + 1]bool
	)
	nums[1] = true
	for i := 2; i*i <= 2*123_456; i++ {
		if nums[i] {
			continue
		}
		for j := i * i; j <= 2*123_456; j += i {
			nums[j] = true
		}
	}

	for {
		fmt.Fscan(r, &n)
		if n == 0 {
			break
		}
		cnt := 0
		for i := n + 1; i <= 2*n; i++ {
			if !nums[i] {
				cnt++
			}
		}
		fmt.Fprintln(w, cnt)
	}
}
