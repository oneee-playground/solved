package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m, q int
		arr     []int
	)

	fmt.Fscan(r, &n)
	arr = make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}
	sort.Ints(arr)

	fmt.Fscan(r, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &q)
		_, found := slices.BinarySearch(arr, q)
		if found {
			fmt.Fprintln(w, 1)
			continue
		}
		fmt.Fprintln(w, 0)
	}
}
