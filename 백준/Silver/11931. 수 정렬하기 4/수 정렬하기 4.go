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
		n   int
		arr []int
	)

	fmt.Fscan(r, &n)
	arr = make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for _, n := range arr {
		fmt.Fprintln(w, n)
	}
}
