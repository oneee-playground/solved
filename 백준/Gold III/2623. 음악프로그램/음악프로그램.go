package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m int

		nodes [][]int
		indeg []int
	)

	fmt.Fscan(r, &n, &m)
	nodes = make([][]int, n+1)
	indeg = make([]int, n+1)

	var cnt, cur, last int
	for ; m > 0; m-- {
		fmt.Fscan(r, &cnt, &last)
		for ; cnt > 1; cnt-- {
			fmt.Fscan(r, &cur)
			nodes[last] = append(nodes[last], cur)
			indeg[cur]++
			last = cur
		}
	}

	seq := make([]int, 0, n)

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		node := q[0]
		seq = append(seq, node)
		for _, n := range nodes[node] {
			indeg[n]--
			if indeg[n] == 0 {
				q = append(q, n)
			}
		}
		q = q[1:]
	}

	if len(seq)+1 < len(nodes) {
		fmt.Fprintln(w, 0)
		return
	}

	for _, s := range seq {
		fmt.Fprintln(w, s)
	}
}
