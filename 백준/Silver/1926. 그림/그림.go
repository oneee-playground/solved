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
		n, m         int
		maxArea, cnt int

		grid [][]int
	)
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	grid = make([][]int, n)
	for i := range grid {
		row := make([]int, m)
		for j := range row {
			fmt.Fscan(r, &row[j])
		}
		grid[i] = row
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				cnt++
				area := bfs(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Fprintln(w, cnt)
	fmt.Fprintln(w, maxArea)
}

var (
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
	q  = make([][2]int, 0)
)

func bfs(grid [][]int, y, x int) (area int) {
	q = append(q, [2]int{y, x})
	grid[y][x] = 2
	for len(q) > 0 {
		area++
		p := q[0]
		for i := range dx {
			ny, nx := p[0]+dy[i], p[1]+dx[i]
			if nx < 0 || ny < 0 || ny == len(grid) || nx == len(grid[0]) {
				continue
			}
			if grid[ny][nx] == 1 {
				q = append(q, [2]int{ny, nx})
				grid[ny][nx] = 2
			}
		}
		q = q[1:]
	}
	q = q[:0]
	return area
}
