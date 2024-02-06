package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type elem struct {
	x, y int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n   int
		arr []elem
	)

	fmt.Fscan(r, &n)
	arr = make([]elem, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i].x, &arr[i].y)
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].y == arr[j].y {
			return arr[i].x < arr[j].x
		}
		return arr[i].y < arr[j].y
	})
	for _, item := range arr {
		fmt.Fprintln(w, item.x, item.y)
	}
}
