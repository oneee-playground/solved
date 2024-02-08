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
		n     int
		ropes []int
	)

	fmt.Fscan(r, &n)
	ropes = make([]int, n)
	for i := range ropes {
		fmt.Fscan(r, &ropes[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ropes)))

	var max int
	for i := 0; i < n; i++ {
		v := ropes[i] * (i + 1)
		if v > max {
			max = v
		}
	}
	fmt.Fprintln(w, max)
}
