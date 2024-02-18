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
		n, m     int
		op, a, b int

		arr []int
	)
	fmt.Fscan(r, &n, &m)
	arr = make([]int, n+1)
	for i := range arr {
		arr[i] = i
	}

	var find func(num int) int
	find = func(num int) int {
		v := arr[num]
		if v == num {
			return v
		}
		arr[num] = find(v)
		return arr[num]
	}

	for ; m > 0; m-- {
		fmt.Fscan(r, &op, &a, &b)

		x, y := find(a), find(b)
		if op == 1 {
			res := "NO"
			if x == y {
				res = "YES"
			}
			fmt.Fprintln(w, res)
			continue
		}
		arr[y] = x
	}
}
