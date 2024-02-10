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
		n, m int
		arr  []int
	)

	fmt.Fscan(r, &m, &n)
	arr = make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	var start, end int = 0, 1e9
	for start < end {
		mid := (start + end + 1) / 2
		if ok(arr, m, mid) {
			start = mid
			continue
		}
		end = mid - 1
	}
	fmt.Fprintln(w, start)
}

func ok(arr []int, m, n int) bool {
	cnt := 0
	for _, num := range arr {
		cnt += num / n
		if cnt >= m {
			return true
		}
	}
	return false
}
