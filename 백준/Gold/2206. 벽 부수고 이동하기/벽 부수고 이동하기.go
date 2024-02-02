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

		dx = [4]int{1, 0, -1, 0}
		dy = [4]int{0, 1, 0, -1}

		grid        [][]byte
		visit       [][]bool
		visitBroken [][]bool
		q           [][4]int // y, x, breachCnt dist
	)
	fmt.Fscanf(r, "%d %d\n", &n, &m)

	grid = make([][]byte, n)
	visit = make([][]bool, n)
	visitBroken = make([][]bool, n)
	for j := range grid {
		row := make([]byte, m)
		fmt.Fscanln(r, &row)
		grid[j] = row
		visit[j] = make([]bool, m)
		visitBroken[j] = make([]bool, m)
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
				if breachCnt > 0 {
					continue
				}
				breachCnt = 1
			}
			if breachCnt == 1 {
				if !visit[ny][nx] && !visitBroken[ny][nx] {
					q = append(q, [4]int{ny, nx, 1, dist})
					visitBroken[ny][nx] = true
				}
				continue
			}
			if !visit[ny][nx] {
				q = append(q, [4]int{ny, nx, 0, dist})
				visit[ny][nx] = true
			}

		}
		q = q[1:]
	}
	fmt.Fprintln(w, -1)
}
