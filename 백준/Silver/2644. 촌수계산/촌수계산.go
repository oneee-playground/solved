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
		n, nEdges int
		nodes     [][]int

		start, target int
	)

	fmt.Fscan(r, &n)
	fmt.Fscan(r, &start, &target)
	fmt.Fscan(r, &nEdges)

	nodes = make([][]int, n+1)
	for i := 0; i < nEdges; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		nodes[a] = append(nodes[a], b)
		nodes[b] = append(nodes[b], a)
	}

	visit := make([]bool, n+1)
	visit[start] = true
	q := [][2]int{{start, 0}}
	for len(q) > 0 {
		v, cnt := q[0][0], q[0][1]
		if v == target {
			fmt.Fprintln(w, cnt)
			return
		}
		for _, node := range nodes[v] {
			if visit[node] {
				continue
			}
			visit[node] = true
			q = append(q, [2]int{node, cnt + 1})
		}
		q = q[1:]
	}
	fmt.Fprintln(w, -1)
}
