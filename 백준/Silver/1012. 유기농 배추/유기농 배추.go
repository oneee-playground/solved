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
		tt        int
		n, m, cnt int

		dx = [4]int{1, 0, -1, 0}
		dy = [4]int{0, 1, 0, -1}

		q [][2]int
	)

	grid := make([][]bool, 2500)
	for i := range grid {
		grid[i] = make([]bool, 2500)
	}

	bfs := func(y, x int) {
		head := q

		grid[y][x] = false
		q = append(q, [2]int{y, x})
		for len(q) > 0 {
			p := q[0]
			for i := range dx {
				ny, nx := p[0]+dy[i], p[1]+dx[i]
				if ny < 0 || nx < 0 || ny == n || nx == m {
					continue
				}
				if grid[ny][nx] {
					q = append(q, [2]int{ny, nx})
					grid[ny][nx] = false
				}
			}
			q = q[1:]
		}
		q = head
	}

	fmt.Fscanf(r, "%d\n", &tt)
	for t := 0; t < tt; t++ {
		fmt.Fscanf(r, "%d %d %d\n", &n, &m, &cnt)

		y, x := 0, 0
		for i := 0; i < cnt; i++ {
			fmt.Fscanf(r, "%d %d\n", &y, &x)
			grid[y][x] = true
		}

		cnt := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] {
					bfs(i, j)
					cnt++
				}
			}
		}
		fmt.Fprintln(w, cnt)
	}

}
