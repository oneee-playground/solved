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
		n     int
		unums []int
		snums []int
	)

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(r, &num)
		if num <= 0 {
			snums = append(snums, num)
			continue
		}
		unums = append(unums, num)
	}

	sum := 0
	sort.Sort(sort.Reverse(sort.IntSlice(unums)))
	sort.Sort(sort.IntSlice(snums))

	for len(snums) > 1 {
		sum += snums[0] * snums[1]
		snums = snums[2:]
	}
	for len(unums) > 1 {
		if unums[0] == 1 || unums[1] == 1 {
			break
		}
		sum += unums[0] * unums[1]
		unums = unums[2:]
	}
	for _, n := range unums {
		sum += n
	}
	for _, n := range snums {
		sum += n
	}
	fmt.Fprintln(w, sum)
}
