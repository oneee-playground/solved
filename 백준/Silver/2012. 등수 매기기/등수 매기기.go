package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var (
		n int
        unhappy uint64
		arr        []int
	)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d", &n)

	arr = make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	rank := 1
	for _, expect := range arr {
		unhappy += uint64(math.Abs(float64(rank - expect)))
		rank++
	}
	fmt.Println(unhappy)
}
