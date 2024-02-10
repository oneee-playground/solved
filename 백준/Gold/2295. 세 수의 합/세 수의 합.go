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
		n int

		arr, arr2 []int
	)

	fmt.Fscan(r, &n)
	arr = make([]int, n)
	arr2 = make([]int, 0, n*n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			arr2 = append(arr2, arr[i]+arr[j])
		}
	}

	sort.Ints(arr2)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			_, found := slices.BinarySearch(arr2, arr[i]-arr[j])
			if found {
				fmt.Fprintln(w, arr[i])
				return
			}
		}
	}
}
