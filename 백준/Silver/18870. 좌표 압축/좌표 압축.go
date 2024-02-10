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
		n   int
		arr []int
	)

	fmt.Fscan(r, &n)
	arr = make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	newArr := make([]int, n)
	copy(newArr, arr)
	sort.Ints(newArr)
	newArr = slices.Compact(newArr)
	for _, num := range arr {
		idx, _ := slices.BinarySearch(newArr, num)
		fmt.Fprintf(w, "%d ", idx)
	}
}
