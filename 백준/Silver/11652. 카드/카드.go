package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n   int
		arr []int64
	)

	fmt.Fscan(r, &n)
	arr = make([]int64, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var last, maxInt int64 = 1<<63 - 1, 0
	cnt, max := 0, 0
	for _, item := range arr {
		if item != last {
			cnt = 0
			last = item
		}
		cnt++
		if cnt > max {
			max = cnt
			maxInt = last
		}
	}

	fmt.Fprintln(w, maxInt)
}
