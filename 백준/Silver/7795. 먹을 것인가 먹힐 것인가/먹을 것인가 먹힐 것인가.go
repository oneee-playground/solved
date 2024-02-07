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
		t, n, m int

		arrA, arrB []int
	)

	fmt.Fscan(r, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(r, &n, &m)

		arrA = make([]int, n)
		for i := range arrA {
			fmt.Fscan(r, &arrA[i])
		}

		arrB = make([]int, m)
		for i := range arrB {
			fmt.Fscan(r, &arrB[i])
		}

		sort.Ints(arrA)
		sort.Ints(arrB)

		var cnt, point int
		for _, cur := range arrA {
			for point < m && cur > arrB[point] {
				point++
			}
			cnt += point
		}
		fmt.Fprintln(w, cnt)
	}
}
