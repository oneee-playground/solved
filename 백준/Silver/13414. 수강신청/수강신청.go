package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type e struct {
	idx int
	num string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m    int
		num     string
		hashmap = make(map[string]int)
		order   []e
	)

	fmt.Fscan(r, &m, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &num)
		hashmap[num] = i
	}

	order = make([]e, 0, len(hashmap))
	for num, idx := range hashmap {
		order = append(order, e{idx, num})
	}

	sort.Slice(order, func(i, j int) bool {
		return order[i].idx < order[j].idx
	})

	for i := 0; i < m && i < len(order); i++ {
		fmt.Fprintln(w, order[i].num)
	}
}
