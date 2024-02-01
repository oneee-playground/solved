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

		q       [][2]int
		picture [][]byte
		visit   [][]bool
	)

	bfs := func(y, x int, isColorBlind bool) {
		head := q

		c := picture[y][x]
		q = append(q, [2]int{y, x})
		visit[y][x] = true

		for len(q) > 0 {
			p := q[0]
			for i := range dx {
				ny, nx := p[0]+dy[i], p[1]+dx[i]
				if ny < 0 || nx < 0 || ny == n || nx == n || visit[ny][nx] {
					continue
				}
				cur := picture[ny][nx]
				if c != cur && (!isColorBlind || c == 'B' || cur == 'B') {
					continue
				}
				q = append(q, [2]int{ny, nx})
				visit[ny][nx] = true
			}
			q = q[1:]
		}
		q = head
	}

	fmt.Fscanf(r, "%d\n", &n)
	picture = make([][]byte, n)
	visit = make([][]bool, n)
	for i := range picture {
		row := make([]byte, n)
		fmt.Fscanf(r, "%s\n", &row)
		picture[i] = row
		visit[i] = make([]bool, n)
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visit[i][j] {
				bfs(i, j, false)
				cnt++
			}
		}
	}

	for i := range visit {
		clear(visit[i])
	}

	colorBlindCnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visit[i][j] {
				bfs(i, j, true)
				colorBlindCnt++
			}
		}
	}
	fmt.Fprintln(w, cnt, colorBlindCnt)

}
