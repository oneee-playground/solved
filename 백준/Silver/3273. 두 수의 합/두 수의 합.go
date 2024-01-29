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
		n, x int
	)
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	fmt.Fscan(r, &x)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	cnt, i, j := 0, 0, n-1
	for i < j {
		v := arr[i] + arr[j]
		if v < x {
			i++
			continue
		}
		if v == x {
			cnt++
		}
		j--
	}
	fmt.Fprintln(w, cnt)
}
