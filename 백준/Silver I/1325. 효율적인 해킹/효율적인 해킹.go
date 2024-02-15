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
		nodes[v] = append(nodes[v], u)
	}

	bfs := func(idx int) (cnt int) {
		q := []int{idx}
		visit[idx] = true

		for len(q) > 0 {
			idx = q[0]
			cnt++
			for _, node := range nodes[idx] {
				if visit[node] {
					continue
				}
				q = append(q, node)
				visit[node] = true
			}
			q = q[1:]
		}
		return
	}

	computers := make([]int, 0)
	maxCnt := 0
	for node := range nodes {
		cnt := bfs(node)
		clear(visit)
		if cnt < maxCnt {
			continue
		}
		if cnt > maxCnt {
			maxCnt = cnt
			computers = computers[:0]
		}
		computers = append(computers, node+1)
	}
	for _, c := range computers {
		fmt.Fprintf(w, "%d ", c)
	}
}
