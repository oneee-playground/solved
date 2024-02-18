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
		n, k, m int
		visit   []bool
		nodes   [][]int
	)

	fmt.Fscan(r, &n, &k, &m)
	visit = make([]bool, n+m+1)
	nodes = make([][]int, n+m+1)

	var tmp int
	for i := 1; i <= m; i++ {
		for j := 0; j < k; j++ {
			fmt.Fscan(r, &tmp)
			nodes[tmp] = append(nodes[tmp], i+n)
			nodes[i+n] = append(nodes[i+n], tmp)
		}
	}

	bfs := func(num int) int {
		q := [][2]int{{num, 1}}
		visit[num] = true

		for len(q) > 0 {
			v := q[0]

			if v[0] == n {
				return v[1]
			}

			for _, node := range nodes[v[0]] {
				if visit[node] {
					continue
				}
				cost := v[1]
				if node <= n {
					cost++
				}
				q = append(q, [2]int{node, cost})
				visit[node] = true
			}
			q = q[1:]
		}
		return -1
	}

	depth := bfs(1)
	fmt.Fprintln(w, depth)
}
