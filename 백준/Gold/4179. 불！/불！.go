package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n, m int
		s    string
		grid [][]int
		q    [][3]int

		head = q

		startY, startX int
	)
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	grid = make([][]int, n)
	for i := range grid {
		row := make([]int, m)
		fmt.Fscanf(r, "%s\n", &s)
		for j, c := range s {
			if c == '#' {
				row[j] = -1
				continue
			}
			if c == 'F' {
				q = append(q, [3]int{i, j, 1})
				row[j] = 1
				continue
			}
			if c == 'J' {
				startY, startX = i, j
			}
			row[j] = 0
		}
		grid[i] = row
	}

	for len(q) > 0 {
		fp := q[0]
		q = q[1:]
		for i := range dx {
			ny, nx := fp[0]+dy[i], fp[1]+dx[i]
			if nx < 0 || ny < 0 || ny == n || nx == m {
				continue
			}
			if grid[ny][nx] == 0 {
				q = append(q, [3]int{ny, nx, fp[2] + 1})
				grid[ny][nx] = fp[2] + 1
			}
		}
	}

	q = append(head, [3]int{startY, startX, 1})
	grid[startY][startX] = -1

	for len(q) > 0 {
		pp := q[0]
		q = q[1:]
		for i := range dx {
			ny, nx := pp[0]+dy[i], pp[1]+dx[i]
			if nx < 0 || ny < 0 || ny == n || nx == m {
				fmt.Fprintln(w, pp[2])
				return
			}
			if eta := grid[ny][nx]; eta == 0 || eta > pp[2]+1 {
				q = append(q, [3]int{ny, nx, pp[2] + 1})
				grid[ny][nx] = -1
			}
		}
	}
	fmt.Fprintln(w, "IMPOSSIBLE")
}
