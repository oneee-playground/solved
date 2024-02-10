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
		n, m       int
		arrA, arrB []int
	)

	fmt.Fscan(r, &n, &m)
	arrA = make([]int, n)
	arrB = make([]int, m)
	for i := range arrA {
		fmt.Fscan(r, &arrA[i])
	}
	sort.Ints(arrA)

	for i := range arrB {
		fmt.Fscan(r, &arrB[i])
	}
	sort.Ints(arrB)

	nums := make([]int, 0)
	for _, num := range arrA {
		_, found := slices.BinarySearch(arrB, num)
		if !found {
			nums = append(nums, num)
		}
	}

	fmt.Fprintln(w, len(nums))
	for _, num := range nums {
		fmt.Fprintf(w, "%d ", num)
	}
}
