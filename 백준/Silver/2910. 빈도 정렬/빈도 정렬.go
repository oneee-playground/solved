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
		n, c int

		counts map[int][2]int
		arr    []int
	)

	fmt.Fscan(r, &n, &c)
	counts = make(map[int][2]int)
	arr = make([]int, n)

	for i := range arr {
		fmt.Fscan(r, &arr[i])
		if v, ok := counts[arr[i]]; ok {
			counts[arr[i]] = [2]int{v[0] + 1, v[1]}
		} else {
			counts[arr[i]] = [2]int{1, i}
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		first, second := counts[arr[i]], counts[arr[j]]
		if first[0] == second[0] {
			return first[1] < second[1]
		}
		return first[0] > second[0]
	})

	for _, num := range arr {
		fmt.Fprintf(w, "%d ", num)
	}
}
