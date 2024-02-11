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
		n      int
		fluids []int
	)

	fmt.Fscan(r, &n)
	fluids = make([]int, n)
	for i := range fluids {
		fmt.Fscan(r, &fluids[i])
	}

	var closest int = 1e10

	start, end := 0, n-1
	for start < end {
		v := fluids[start] + fluids[end]
		if abs(v) < abs(closest) {
			closest = v
			if v == 0 {
				break
			}
		}
		if v < 0 {
			start++
		} else {
			end--
		}
	}
	fmt.Fprintln(w, closest)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
