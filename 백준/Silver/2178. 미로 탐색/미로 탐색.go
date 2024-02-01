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
		grid [][]byte
	)
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	grid = make([][]byte, n)
	for i := range grid {
		row := make([]byte, m)
		fmt.Fscanf(r, "%s\n", &row)
		grid[i] = row
	}
	fmt.Fprintln(w, bfs(grid, 0, 0))
}

var (
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
	q  = make([][3]int, 0)
)

func bfs(grid [][]byte, y, x int) (min int) {
	q = append(q, [3]int{y, x, 1})
	grid[y][x] = 'V'
	for len(q) > 0 {
		p := q[0]
		if p[0] == len(grid)-1 && p[1] == len(grid[0])-1 {
			return p[2]
		}

		for i := range dx {
			ny, nx := p[0]+dy[i], p[1]+dx[i]
			if nx < 0 || ny < 0 || ny == len(grid) || nx == len(grid[0]) {
				continue
			}
			if grid[ny][nx] == '1' {
				q = append(q, [3]int{ny, nx, p[2] + 1})
				grid[ny][nx] = 'V'
			}
		}
		q = q[1:]
	}
	q = q[:0]
	return
}
