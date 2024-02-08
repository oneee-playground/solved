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
		n    int
		a, b []int
	)

	fmt.Fscan(r, &n)
	a, b = make([]int, n), make([]int, n)

	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	for i := range b {
		fmt.Fscan(r, &b[i])
	}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	sum := 0
	for i := range a {
		sum += a[i] * b[i]
	}
	fmt.Fprintln(w, sum)
}
