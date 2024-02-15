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

	cnt := 0
	bfs := func(idx int) {
		q := [][2]int{{idx, 0}}
		visit[idx] = true
		for len(q) > 0 {
			cnt++
			i := q[0]
			if i[1] == 2 {
				q = q[1:]
				continue
			}
			for _, node := range nodes[i[0]] {
				if visit[node] {
					continue
				}
				q = append(q, [2]int{node, i[1] + 1})
				visit[node] = true
			}
			q = q[1:]
		}
	}

	bfs(0)
	fmt.Fprintln(w, cnt-1)
}
