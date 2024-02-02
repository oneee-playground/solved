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
		n int

		dx = [4]int{1, -1, 0, 0}
		dy = [4]int{0, 0, 1, -1}

		grid  [][]int
		visit [][]bool
		isle  [][]int

		q1 [][2]int
		q2 [][3]int
	)

	bfs := func(y, x, cnt int) int {
		head1, head2 := q1, q2
		defer func() { q1 = head1; q2 = head2 }()

		q1 = append(q1, [2]int{y, x})
		isle[y][x] = cnt

		for len(q1) > 0 {
			p := q1[0]
			for i := range dx {
				ny, nx := p[0]+dy[i], p[1]+dx[i]
				if ny < 0 || nx < 0 || ny >= n || nx >= n {
					continue
				}
				if grid[ny][nx] == 0 && !visit[ny][nx] {
					q2 = append(q2, [3]int{ny, nx, 1})
					visit[ny][nx] = true
					continue
				}
				if grid[ny][nx] == 1 && isle[ny][nx] == 0 {
					q1 = append(q1, [2]int{ny, nx})
					isle[ny][nx] = cnt
				}

			}
			q1 = q1[1:]
		}

		for len(q2) > 0 {
			p := q2[0]
			for i := range dx {
				ny, nx := p[0]+dy[i], p[1]+dx[i]
				if ny < 0 || nx < 0 || ny >= n || nx >= n {
					continue
				}
				if grid[ny][nx] == 1 {
					if isle[ny][nx] != cnt {
						return p[2]
					}
					continue
				}
				if !visit[ny][nx] {
					q2 = append(q2, [3]int{ny, nx, p[2] + 1})
					visit[ny][nx] = true
				}
			}
			q2 = q2[1:]
		}
		panic("wtf")
	}

	fmt.Fscan(r, &n)
	grid = make([][]int, n)
	visit = make([][]bool, n)
	isle = make([][]int, n)
	for i := range grid {
		row := make([]int, n)
		for j := range row {
			fmt.Fscan(r, &row[j])
		}
		grid[i] = row
		visit[i] = make([]bool, n)
		isle[i] = make([]int, n)
	}

	cnt, minLen := 1, n*n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && isle[i][j] == 0 {
				l := bfs(i, j, cnt)
				if l < minLen {
					minLen = l
				}
				cnt++

				for k := range visit {
					clear(visit[k])
				}
			}
		}
	}
	fmt.Fprintln(w, minLen)
}
