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
		k, n int

		arr []uint64

		start, end uint64
	)

	fmt.Fscan(r, &k, &n)
	arr = make([]uint64, k)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	start, end = 0, 1<<31-1
	for start < end {
		mid := (start + end + 1) / 2
		if ok(arr, n, mid) {
			start = mid
			continue
		}
		end = mid - 1
	}
	fmt.Fprintln(w, start)
}

func ok(arr []uint64, n int, l uint64) bool {
	cnt := 0
	for _, num := range arr {
		cnt += int(num / l)
	}
	return cnt >= n
}
