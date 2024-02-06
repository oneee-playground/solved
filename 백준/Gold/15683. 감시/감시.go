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

	fmt.Fscan(r, &n, &m)
	grid = make([][]int, n)
	check = make([][]bool, n)
	for i := range grid {
		row := make([]int, m)
		for j := range row {
			fmt.Fscan(r, &row[j])
			if row[j] == 6 {
				wallCnt++
			}
			if row[j] >= 1 && row[j] <= 5 {
				cctvs = append(cctvs, [4]int{i, j, 0, row[j]})
			}
		}
		grid[i] = row
		check[i] = make([]bool, m)
	}
	min = m * n
	find(0)
	fmt.Fprintln(w, min-wallCnt)
}

var (
	n, m int

	grid  [][]int
	cctvs [][4]int // y, x, dir, num
	check [][]bool

	min, wallCnt int

	dirs = [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func observe(y, x, dy, dx int) {
	ny, nx := y, x
	for {
		ny, nx = ny+dy, nx+dx
		if ny < 0 || nx < 0 || ny >= n || nx >= m {
			return
		}
		if grid[ny][nx] == 6 {
			return
		}
		check[ny][nx] = true
	}
}

func getCnt() (cnt int) {
	for _, cctv := range cctvs {
		y, x := cctv[0], cctv[1]
		check[y][x] = true

		num := cctv[3]

		dir := cctv[2]
		dy, dx := dirs[dir][0], dirs[dir][1]
		switch num {
		case 1:
			observe(y, x, dy, dx)
		case 2:
			observe(y, x, dy, dx)
			observe(y, x, -dy, -dx)
		case 3:
			observe(y, x, dy, dx)
			observe(y, x, -dx, dy)
		case 4:
			observe(y, x, dy, dx)
			observe(y, x, -dy, -dx)
			observe(y, x, -dx, dy)
		case 5:
			observe(y, x, 1, 0)
			observe(y, x, 0, 1)
			observe(y, x, 0, -1)
			observe(y, x, -1, 0)
		}
	}

	for i := range check {
		for _, ok := range check[i] {
			if !ok {
				cnt++
			}
		}
	}

	for i := range check {
		clear(check[i])
	}
	return
}

func find(i int) {
	if i == len(cctvs) {
		cnt := getCnt()
		if cnt < min {
			min = cnt
		}
		return
	}
	if cctvs[i][3] == 5 {
		find(i + 1)
		return
	}

	for j := 0; j < 4; j++ {
		cctvs[i][2] = j
		find(i + 1)
	}
}
