package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m   int
	checks [][]bool
	grid   [][]int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &m)
	grid = make([][]int, n)
	checks = make([][]bool, n)
	for i := 0; i < n; i++ {
		row := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &row[j])
		}
		grid[i] = row
		checks[i] = make([]bool, m)
	}

	year := 0
	for {
		ice := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] > 0 && !checks[i][j] {
					ice++
					dfs(i, j)
				}
			}
		}
		if ice != 1 {
			if ice == 0 {
				year = 0
			}
			break
		}
		for i := 0; i < n; i++ {
			clear(checks[i])
		}
		year++
	}
	fmt.Fprintln(w, year)
}

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

func dfs(y, x int) {
	checks[y][x] = true
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if checks[ny][nx] {
			continue
		}
		if grid[ny][nx] > 0 {
			dfs(ny, nx)
			continue
		}
		grid[y][x]--
	}
}
