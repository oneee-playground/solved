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
		m, n, h   int
		tomatoCnt int
		days      int

		dx = [6]int{1, 0, -1, 0, 0, 0}
		dy = [6]int{0, 1, 0, -1, 0, 0}
		dz = [6]int{0, 0, 0, 0, 1, -1}

		q   [][4]int
		box [][][]int
	)
	fmt.Fscanf(r, "%d %d %d\n", &m, &n, &h)
	box = make([][][]int, h)
	for i := range box {
		grid := make([][]int, n)
		for j := range grid {
			row := make([]int, m)
			for k := range row {
				fmt.Fscan(r, &row[k])
				if row[k] == 0 {
					tomatoCnt++
				}
				if row[k] == 1 {
					q = append(q, [4]int{i, j, k, 0})
				}
			}
			grid[j] = row
		}
		box[i] = grid
	}

	for len(q) > 0 {
		p := q[0]
		days = p[3]
		for i := range dx {
			nh, ny, nx := p[0]+dz[i], p[1]+dy[i], p[2]+dx[i]
			if nh < 0 || ny < 0 || nx < 0 || nh == h || ny == n || nx == m {
				continue
			}
			if box[nh][ny][nx] == 0 {
				q = append(q, [4]int{nh, ny, nx, days + 1})
				box[nh][ny][nx] = 1
				tomatoCnt--
			}
		}
		q = q[1:]
	}

	if tomatoCnt > 0 {
		days = -1
	}

	fmt.Fprintln(w, days)
}
