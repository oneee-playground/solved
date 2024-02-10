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
	compressed := compress(newArr)

	for _, num := range arr {
		idx, _ := slices.BinarySearch(compressed, num)
		fmt.Fprintf(w, "%d ", idx)
	}
}

func compress(arr []int) []int {
	newArr := make([]int, 0)
	last := 0
	for _, num := range arr {
		if num != last {
			newArr = append(newArr, num)
			last = num
		}
	}
	return newArr
}
