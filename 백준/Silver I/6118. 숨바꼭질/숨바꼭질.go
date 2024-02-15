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

	bfs := func(idx int) ([]int, int) {
		q := [][2]int{{idx, 0}}
		visit[idx] = true

		last := 0
		candidates := make([]int, 0)
		for len(q) > 0 {
			i := q[0]
			if i[1] > last {
				candidates = candidates[:0]
			}
			candidates = append(candidates, i[0])
			for _, node := range nodes[i[0]] {
				if visit[node] {
					continue
				}
				q = append(q, [2]int{node, i[1] + 1})
				visit[node] = true
			}
			last = i[1]
			q = q[1:]
		}
		return candidates, last
	}

	items, dist := bfs(0)
	sort.Ints(items)
	fmt.Fprintln(w, items[0]+1, dist, len(items))
}
