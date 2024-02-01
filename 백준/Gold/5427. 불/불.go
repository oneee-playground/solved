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
		tt   int
		n, m int
		cnt  int

		dx = [4]int{1, -1, 0, 0}
		dy = [4]int{0, 0, 1, -1}

		grid [][]byte
		qp   [][3]int
		qf   [][2]int
	)
	fmt.Fscanf(r, "%d\n", &tt)
	for t := 0; t < tt; t++ {
		pHead, fHead := qp, qf

		fmt.Fscanf(r, "%d %d\n", &m, &n)
		grid = make([][]byte, n)
		for i := range grid {
			row := make([]byte, m)
			fmt.Fscanf(r, "%s\n", &row)
			for j := range row {
				if row[j] == '@' {
					qp = append(qp, [3]int{i, j, 0})
				}
				if row[j] == '*' {
					qf = append(qf, [2]int{i, j})
				}
			}
			grid[i] = row
		}

		for len(qp) > 0 {
			pLen, fLen := len(qp), len(qf)

			for i := 0; i < fLen; i++ {
				p := qf[0]
				qf = qf[1:]
				for j := range dx {
					ny, nx := p[0]+dy[j], p[1]+dx[j]
					if ny < 0 || nx < 0 || ny >= n || nx >= m {
						continue
					}
					if grid[ny][nx] == '.' || grid[ny][nx] == '@' {
						qf = append(qf, [2]int{ny, nx})
						grid[ny][nx] = '*'
					}
				}
			}

			for i := 0; i < pLen; i++ {
				p := qp[0]
				qp = qp[1:]
				for j := range dx {
					ny, nx := p[0]+dy[j], p[1]+dx[j]
					if ny < 0 || nx < 0 || ny >= n || nx >= m {
						cnt = p[2] + 1
						i = pLen
						qp = nil
						break
					}
					if grid[ny][nx] == '.' {
						qp = append(qp, [3]int{ny, nx, p[2] + 1})
						grid[ny][nx] = '@'
					}
				}
			}
		}
		if cnt == 0 {
			fmt.Fprintln(w, "IMPOSSIBLE")
		} else {
			fmt.Fprintln(w, cnt)
		}

		qp, qf = pHead, fHead
		cnt = 0
	}
}
