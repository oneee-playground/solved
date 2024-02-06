package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type elem struct {
	name string
	age  int
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
		fmt.Fscan(r, &arr[i].age, &arr[i].name)
	}

	slices.SortStableFunc(arr, func(a, b elem) int {
		return a.age - b.age
	})
	for _, item := range arr {
		fmt.Fprintln(w, item.age, item.name)
	}
}
