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

		dx = [4]int{1, 0, -1, 0}
		dy = [4]int{0, 1, 0, -1}

		grid  [][]int
		visit [][]bool
		q     [][2]int
	)
	fmt.Fscanf(r, "%d\n", &n)
	grid = make([][]int, n)
	visit = make([][]bool, n)
	for i := range grid {
		row := make([]int, n)
		for j := range row {
			fmt.Fscan(r, &row[j])
		}
		grid[i] = row
		visit[i] = make([]bool, n)
	}

	bfs := func(h, y, x int) {
		head := q
		q = append(q, [2]int{y, x})
		visit[y][x] = true

		for len(q) > 0 {
			p := q[0]
			for i := range dx {
				ny, nx := p[0]+dy[i], p[1]+dx[i]
				if ny < 0 || nx < 0 || ny >= n || nx >= n {
					continue
				}
				if grid[ny][nx] > h && !visit[ny][nx] {
					q = append(q, [2]int{ny, nx})
					visit[ny][nx] = true
				}
			}
			q = q[1:]
		}
		q = head
	}

	var maxCnt, curCnt int = 0, 1
	for h := 0; curCnt > 0; h++ {
		curCnt = 0

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] > h && !visit[i][j] {
					curCnt++
					bfs(h, i, j)
				}
			}
		}

		if curCnt > maxCnt {
			maxCnt = curCnt
		}

		for i := range visit {
			clear(visit[i])
		}
	}
	fmt.Fprintln(w, maxCnt)
}
