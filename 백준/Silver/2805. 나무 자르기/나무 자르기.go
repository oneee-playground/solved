package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n      int
		m, max uint64
		trees  []uint64
	)

	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d %d", &n, &m)
	trees = make([]uint64, n)
	for i := 0; i < n; i++ {
		var v uint64
		fmt.Fscan(r, &v)
		if v > max {
			max = v
		}
		trees[i] = v
	}

	start, end := uint64(0), max
	for start+1 < end {
		got := uint64(0)
		mid := (start + end) / 2
		for i := 0; i < len(trees); i++ {
			if h := trees[i]; mid < h {
				got += h - mid
			}
		}
		if got >= m {
			start = mid
		} else {
			end = mid
		}
	}
	fmt.Println(start)
}
