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
		visit []bool
	)
	fmt.Fscan(r, &n, &m)
	nodes = make([][]int, n)
	visit = make([]bool, n)

	var u, v int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &u, &v)
		u--
		v--
		nodes[u] = append(nodes[u], v)
		nodes[v] = append(nodes[v], u)
	}

	bfs := func(idx int) {
		q := []int{idx}
		visit[idx] = true
		for len(q) > 0 {
			i := q[0]
			for _, node := range nodes[i] {
				if visit[node] {
					continue
				}
				q = append(q, node)
				visit[node] = true
			}
			q = q[1:]
		}
	}

	cnt := 0
	for idx := range nodes {
		if visit[idx] {
			continue
		}
		cnt++
		bfs(idx)
	}
	fmt.Fprintln(w, cnt)
}
