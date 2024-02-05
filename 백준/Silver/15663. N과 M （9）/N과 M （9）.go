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
	check = make([]bool, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}
	sort.Ints(nums)

	result = make([]int, m)
	find(w, 0)
}

var (
	n, m   int
	nums   []int
	result []int
	check  []bool
)

func find(w io.Writer, i int) {
	if i == m {
		for _, num := range result {
			fmt.Fprintf(w, "%d ", num)
		}
		fmt.Fprintln(w)
		return
	}

	last := 0
	for j := 0; j < n; j++ {
		if check[j] || nums[j] == last {
			continue
		}

		result[i] = nums[j]
		check[j] = true
		find(w, i+1)
		check[j] = false

		last = nums[j]
	}
}
