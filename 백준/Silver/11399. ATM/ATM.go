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
		n      int
		people []int
	)
	fmt.Fscan(r, &n)

	people = make([]int, n)
	for i := range people {
		fmt.Fscan(r, &people[i])
	}

	sort.Ints(people)
	sum, wait := 0, 0
	for _, t := range people {
		wait += t
		sum += wait
	}
	fmt.Fprintln(w, sum)
}
