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

		grid    [][]byte
		hashmap = make(map[string]uint)

		dx = [8]int{0, 0, 1, -1, 1, -1, 1, -1}
		dy = [8]int{1, -1, 0, 0, -1, 1, 1, -1}

		str = make([]byte, 5)
	)

	fmt.Fscan(r, &n, &m, &k)
	grid = make([][]byte, n)
	for i := range grid {
		row := make([]byte, m)
		fmt.Fscan(r, &row)
		grid[i] = row
	}

	var dfs func(y, x, cnt int)
	dfs = func(y, x, cnt int) {
		str[cnt-1] = grid[y][x]
		hashmap[string(str[:cnt])]++
		if cnt == len(str) {
			return
		}
		for i := range dx {
			ny, nx := adjust(y+dy[i], n), adjust(x+dx[i], m)
			dfs(ny, nx, cnt+1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j, 1)
		}
	}

	var query string
	for ; k > 0; k-- {
		fmt.Fscan(r, &query)
		fmt.Fprintln(w, hashmap[query])
	}
}

func adjust(cur, edge int) int {
	if cur < 0 {
		return edge - 1
	}
	if cur == edge {
		return 0
	}
	return cur
}
