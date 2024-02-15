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
		t, n, m int

		nodes  [][]int
		colors []int

		dfs func(idx, color int) bool
	)

	dfs = func(idx, color int) bool {
		colors[idx] = color
		for _, node := range nodes[idx] {
			if node == idx || colors[node] == color {
				return false
			}
			if colors[node] == 0 {
				if ok := dfs(node, -color); !ok {
					return false
				}
			}
		}
		return true
	}

	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &n, &m)
		nodes = make([][]int, n)
		colors = make([]int, n)

		var u, v int
		for i := 0; i < m; i++ {
			fmt.Fscan(r, &u, &v)
			u--
			v--
			nodes[u] = append(nodes[u], v)
			nodes[v] = append(nodes[v], u)
		}

		res := "YES"
		for i := 0; i < n; i++ {
			if colors[i] != 0 {
				continue
			}
			if ok := dfs(i, 1); !ok {
				res = "NO"
				break
			}
		}
		fmt.Fprintln(w, res)
	}
}
