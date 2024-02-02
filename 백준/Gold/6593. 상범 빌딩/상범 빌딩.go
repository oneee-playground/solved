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
		h, n, m int

		dx = [6]int{1, 0, -1, 0, 0, 0}
		dy = [6]int{0, 1, 0, -1, 0, 0}
		dz = [6]int{0, 0, 0, 0, 1, -1}

		building [][][]byte
		q        [][4]int
	)
	for {
		fmt.Fscanf(r, "%d %d %d\n", &h, &n, &m)
		if h == 0 && n == 0 && m == 0 {
			break
		}
		head := q

		building = make([][][]byte, h)
		for i := range building {
			grid := make([][]byte, n)
			for j := range grid {
				row := make([]byte, m)
				fmt.Fscanln(r, &row)
				for k := range row {
					if row[k] == 'S' {
						q = append(q, [4]int{i, j, k, 0})
					}
				}
				grid[j] = row
			}
			building[i] = grid
			fmt.Fscanln(r)
		}

		cnt := -1
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for i := range dx {
				nz, ny, nx := p[0]+dz[i], p[1]+dy[i], p[2]+dx[i]
				if nz < 0 || ny < 0 || nx < 0 || nz >= h || ny >= n || nx >= m {
					continue
				}
				if building[nz][ny][nx] == 'E' {
					cnt = p[3] + 1
					q = nil
					break
				}
				if building[nz][ny][nx] == '.' {
					q = append(q, [4]int{nz, ny, nx, p[3] + 1})
					building[nz][ny][nx] = 'S'
				}
			}
		}
		q = head

		if cnt != -1 {
			fmt.Fprintf(w, "Escaped in %d minute(s).\n", cnt)
			continue
		}
		fmt.Fprintln(w, "Trapped!")
	}
}
