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
		n, m, v int

		nodes [][]int
		visit []bool
	)
	fmt.Fscan(r, &n, &m, &v)
	nodes = make([][]int, n)
	visit = make([]bool, n)
	v--

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b)
		a--
		b--
		nodes[a] = append(nodes[a], b)
		nodes[b] = append(nodes[b], a)
	}

	for _, node := range nodes {
		sort.Ints(node)
	}

	bfs := func(idx int) {
		q := []int{idx}
		visit[idx] = true
		for len(q) > 0 {
			i := q[0]
			fmt.Fprintf(w, "%d ", i+1)
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

	var dfs func(idx int)
	dfs = func(idx int) {
		visit[idx] = true
		fmt.Fprintf(w, "%d ", idx+1)
		for _, node := range nodes[idx] {
			if visit[node] {
				continue
			}
			dfs(node)
		}
	}

	dfs(v)
	fmt.Fprintln(w)
	clear(visit)
	bfs(v)
}
