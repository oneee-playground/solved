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
		n int

		nodes [][]int
		visit []bool
	)
	fmt.Fscan(r, &n)
	nodes = make([][]int, n)
	visit = make([]bool, n)

	var u, v int
	for {
		fmt.Fscan(r, &u, &v)
		if u == -1 && v == -1 {
			break
		}
		u--
		v--
		nodes[u] = append(nodes[u], v)
		nodes[v] = append(nodes[v], u)
	}

	bfs := func(idx int) (depth int) {
		q := [][2]int{{idx, 0}}
		visit[idx] = true
		for len(q) > 0 {
			idx, depth = q[0][0], q[0][1]
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

	minDepth := n
	candidates := make([]int, 0)
	for node := range nodes {
		depth := bfs(node)
		clear(visit)
		if depth > minDepth {
			continue
		}
		if depth < minDepth {
			minDepth = depth
			candidates = candidates[:0]
		}
		candidates = append(candidates, node+1)
	}

	fmt.Fprintln(w, minDepth, len(candidates))
	for _, c := range candidates {
		fmt.Fprintf(w, "%d ", c)
	}
}
