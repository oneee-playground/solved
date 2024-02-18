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
		n, m, k int
		cntP    int
		p       []int
		nodes   [][]int
	)

	fmt.Fscan(r, &n, &m, &cntP)

	p = make([]int, n+1)
	for i := range p {
		p[i] = i
	}

	for i := 0; i < cntP; i++ {
		fmt.Fscan(r, &k)
		p[k] = 0
	}

	var find func(i int) int
	find = func(i int) int {
		if i == p[i] {
			return i
		}
		p[i] = find(p[i])
		return p[i]
	}

	merge := func(a, b int) {
		x, y := find(a), find(b)
		if x < y {
			p[y] = p[x]
		} else {
			p[x] = p[y]
		}
	}

	nodes = make([][]int, m)
	for i := 0; i < m; i++ {
		var last int
		fmt.Fscan(r, &cntP, &last)
		nodes[i] = append(nodes[i], last)
		for j := 1; j < cntP; j++ {
			fmt.Fscan(r, &k)
			nodes[i] = append(nodes[i], k)
			merge(last, k)
			last = k
		}
	}

	cnt := m
	for _, node := range nodes {
		for _, person := range node {
			if find(person) == 0 {
				cnt--
				break
			}
		}
	}
	fmt.Fprintln(w, cnt)
}
