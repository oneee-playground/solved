package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for {
		fmt.Fscan(r, &n)
		if n == 0 {
			return
		}
		nums = make([]int, n)
		for i := range nums {
			fmt.Fscan(r, &nums[i])
		}

		find(w, 0)
		fmt.Fprintln(w)
	}
}

var (
	n    int
	nums []int
	idxs [6]int
)

func find(w io.Writer, i int) {
	if i == 6 {
		for _, idx := range idxs {
			fmt.Fprintf(w, "%d ", nums[idx])
		}
		fmt.Fprintln(w)
		return
	}

	j := 0
	if i > 0 {
		j = idxs[i-1] + 1
	}
	for ; j < n; j++ {
		idxs[i] = j
		find(w, i+1)
	}
}
