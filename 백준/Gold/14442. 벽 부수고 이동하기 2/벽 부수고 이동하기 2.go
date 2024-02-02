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

		dx = [4]int{1, 0, -1, 0}
		dy = [4]int{0, 1, 0, -1}

		grid  [][]byte
		visit [][][]bool
		q     [][4]int // y, x, breachCnt dist
	)
	fmt.Fscanf(r, "%d %d %d\n", &n, &m, &k)

	grid = make([][]byte, n)
	visit = make([][][]bool, k+1)
	for i := range visit {
		visit[i] = make([][]bool, n)
		for j := range visit[i] {
			visit[i][j] = make([]bool, m)
		}
	}
	for i := range grid {
		row := make([]byte, m)
		fmt.Fscanln(r, &row)
		grid[i] = row
	}

	q = append(q, [4]int{0, 0, 0, 1})

	for len(q) > 0 {
		p := q[0]
		if p[0] == n-1 && p[1] == m-1 {
			fmt.Fprintln(w, p[3])
			return
		}
		for i := range dx {
			ny, nx := p[0]+dy[i], p[1]+dx[i]
			if ny < 0 || nx < 0 || ny >= n || nx >= m {
				continue
			}
			breachCnt, dist := p[2], p[3]+1
			if grid[ny][nx] == '1' {
				if breachCnt == k {
					continue
				}
				breachCnt++
			}

			visited := false
			for j := 0; j <= breachCnt; j++ {
				if visit[j][ny][nx] {
					visited = true
					break
				}
			}
			if !visited {
				q = append(q, [4]int{ny, nx, breachCnt, dist})
				visit[breachCnt][ny][nx] = true
			}
		}
		q = q[1:]
	}
	fmt.Fprintln(w, -1)
}
