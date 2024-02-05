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

	fmt.Fscan(r, &n, &m)
	nums = make([]int, m)
	find(w, 0)
}

var (
	n, m int
	nums []int
)

func find(w io.Writer, i int) {
	if i == m {
		for _, num := range nums {
			fmt.Fprintf(w, "%d ", num)
		}
		fmt.Fprintln(w)
		return
	}

	for j := 1; j <= n; j++ {
		nums[i] = j
		find(w, i+1)
	}
}
