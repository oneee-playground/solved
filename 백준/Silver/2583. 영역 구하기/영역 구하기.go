package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		m, n, k int

		board [][]int
	)
	fmt.Fscanf(r, "%d %d %d\n", &m, &n, &k)
	board = make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
	}

	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscanf(r, "%d %d %d %d\n", &x1, &y1, &x2, &y2)
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				board[y][x] = 1
			}
		}
	}

	cnt := 0
	areas := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				cnt++
				areas = append(areas, bfs(board, i, j))
			}
		}
	}
	fmt.Fprintln(w, cnt)
	sort.Slice(areas, func(i, j int) bool {
		return areas[i] < areas[j]
	})
	for i := 0; i < len(areas); i++ {
		fmt.Fprintf(w, "%d ", areas[i])
	}
}

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}

	q [][2]int
)

func bfs(board [][]int, y, x int) (area int) {
	head := q
	q = append(q, [2]int{y, x})
	board[y][x] = 1

	for len(q) > 0 {
		area++
		p := q[0]
		for i := range dx {
			ny, nx := p[0]+dy[i], p[1]+dx[i]
			if ny < 0 || nx < 0 || ny >= len(board) || nx >= len(board[0]) {
				continue
			}
			if board[ny][nx] == 0 {
				q = append(q, [2]int{ny, nx})
				board[ny][nx] = 1
			}
		}
		q = q[1:]
	}
	q = head
	return
}
