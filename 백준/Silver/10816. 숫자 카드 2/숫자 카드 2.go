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
		cnt := upperIdx(arr, q) - lowerIdx(arr, q)
		fmt.Fprintf(w, "%d ", cnt)
	}
}

func lowerIdx(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] >= target {
			end = mid
			continue
		}
		start = mid + 1
	}
	return start
}

func upperIdx(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid
			continue
		}
		start = mid + 1
	}
	return start
}
