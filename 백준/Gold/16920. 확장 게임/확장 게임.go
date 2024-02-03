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
		n, m, p int
		players [][2]int

		grid [][]byte

		qs [][][3]int

		dx = [4]int{1, -1, 0, 0}
		dy = [4]int{0, 0, 1, -1}
	)

	fmt.Fscan(r, &n, &m, &p)

	players = make([][2]int, p)
	qs = make([][][3]int, p)
	for i := range players {
		fmt.Fscan(r, &players[i][0])
		qs[i] = make([][3]int, 0)
	}

	grid = make([][]byte, n)
	for i := range grid {
		row := make([]byte, m)
		fmt.Fscan(r, &row)
		grid[i] = row
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if v := grid[i][j]; v != '.' && v != '#' {
				idx := v - '1'
				qs[idx] = append(qs[idx], [3]int{i, j, 0})
				players[idx][1]++
			}
		}
	}

	for {
		proceed := false
		for idx, q := range qs {
			for len(q) > 0 {
				if q[0][2] == players[idx][0] {
					for i := range q {
						q[i][2] = 0
					}
					qs[idx] = q
					break
				}
				size := len(q)
				for i := 0; i < size; i++ {
					p := q[0]
					dist := p[2] + 1
					for i := range dx {
						ny, nx := p[0]+dy[i], p[1]+dx[i]
						if ny < 0 || nx < 0 || ny >= n || nx >= m {
							continue
						}
						if grid[ny][nx] == '.' {
							q = append(q, [3]int{ny, nx, dist})
							grid[ny][nx] = byte(idx) + '1'
							players[idx][1]++
							proceed = true
						}
					}
					q = q[1:]
				}
			}
		}
		if !proceed {
			break
		}
	}
	for _, p := range players {
		fmt.Fprintf(w, "%d ", p[1])
	}
}
