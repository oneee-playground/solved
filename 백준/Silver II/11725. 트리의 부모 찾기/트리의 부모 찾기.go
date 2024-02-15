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
	)

	fmt.Fscan(r, &n)
	nodes = make([][]int, n+1)

	var a, b int
	for i := 1; i < n; i++ {
		fmt.Fscan(r, &a, &b)
		nodes[a] = append(nodes[a], b)
		nodes[b] = append(nodes[b], a)
	}

	p := make([]int, n+1)
	visit := make([]bool, n+1)

	var dfs func(idx int)
	dfs = func(idx int) {
		visit[idx] = true
		for _, node := range nodes[idx] {
			if visit[node] {
				continue
			}
			p[node] = idx
			dfs(node)
		}
	}

	dfs(1)
	for _, parent := range p[2:] {
		fmt.Fprintln(w, parent)
	}
}
