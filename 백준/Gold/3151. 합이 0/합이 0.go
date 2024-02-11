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
		n        int
		students []int
	)

	fmt.Fscan(r, &n)
	students = make([]int, n)
	for i := range students {
		fmt.Fscan(r, &students[i])
	}
	sort.Ints(students)

	cnt := uint64(0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			v := -(students[i] + students[j])

			lowerIdx := func(cur, t int) int { return cur - t }
			upperIdx := func(cur, t int) int { return cur - t - 1 }

			idx1, _ := slices.BinarySearchFunc(students[j+1:], v, lowerIdx)
			idx2, _ := slices.BinarySearchFunc(students[j+1:], v, upperIdx)
			cnt += uint64(idx2 - idx1)
		}
	}
	fmt.Fprintln(w, cnt)
}
