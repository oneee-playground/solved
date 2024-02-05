package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &m)
	nums = make([]int, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}
	sort.Ints(nums)

	idxs = make([]int, m)
	find(w, 0)
}

var (
	n, m int
	nums []int
	idxs []int
)

func find(w io.Writer, i int) {
	if i == m {
		for _, idx := range idxs {
			fmt.Fprintf(w, "%d ", nums[idx])
		}
		fmt.Fprintln(w)
		return
	}

	for j := 0; j < n; j++ {
		idxs[i] = j
		find(w, i+1)
	}
}
