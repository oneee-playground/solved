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
		if u == -1 && v == -1 {
			break
		}
		u--
		v--
		nodes[u] = append(nodes[u], v)
		nodes[v] = append(nodes[v], u)
	}

	bfs := func(idx int) (sum int) {
		q := [][2]int{{idx, 0}}
		visit[idx] = true

		sum = 0
		for len(q) > 0 {
			idx, depth := q[0][0], q[0][1]
			sum += depth
			for _, node := range nodes[idx] {
				if visit[node] {
					continue
				}
				q = append(q, [2]int{node, depth + 1})
				visit[node] = true
			}
			q = q[1:]
		}
		return
	}

	minSum, minIdx := n*m, 0
	for node := range nodes {
		sum := bfs(node)
		clear(visit)
		if sum < minSum {
			minSum = sum
			minIdx = node
		}
	}
	fmt.Fprintln(w, minIdx+1)
}
