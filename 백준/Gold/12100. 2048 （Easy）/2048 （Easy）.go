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

	fmt.Fscan(r, &n)
	for i := 0; i <= 5; i++ {
		states[i] = make([][]int, n)
	}

	checks = make([][]bool, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &row[j])

		}
		checks[i] = make([]bool, n)
		states[0][i] = row
	}
	for i := 1; i <= 5; i++ {
		for j := 0; j < n; j++ {
			states[i][j] = make([]int, n)
		}
	}

	simulate(1)
	fmt.Fprintln(w, theMax)
}

var (
	n      int
	theMax int

	states [6][][]int

	checks [][]bool

	dy = [4]int{1, -1, 0, 0}
	dx = [4]int{0, 0, 1, -1}
)

func push(state, idx int) {
	start := 0
	if idx%2 == 1 {
		start = n - 1
	}
	now := states[state]

	y, x := start, start
	// 행 반복
	for y >= 0 && x >= 0 && y < n && x < n {
		ny, nx := y, x
		// 열 반복
		for ny >= 0 && nx >= 0 && ny < n && nx < n {
			if now[ny][nx] != 0 {
				py, px := ny, nx
				for {
					py, px = py-dy[idx], px-dx[idx]
					if py >= 0 && px >= 0 && py < n && px < n {
						if now[py][px] == 0 {
							continue
						}
						if now[py][px] == now[ny][nx] && !checks[py][px] {
							now[py][px] *= 2
							now[ny][nx] = 0
							checks[py][px] = true
							break
						}
					}
					py, px = py+dy[idx], px+dx[idx]
					v := now[ny][nx]
					now[ny][nx] = 0
					now[py][px] = v
					break
				}
			}
			ny, nx = ny+dy[idx], nx+dx[idx]
		}
		y, x = y+dx[idx], x+dy[idx]
	}

	for i := range checks {
		clear(checks[i])
	}
}

func simulate(i int) {
	if i == 6 {
		for _, row := range states[5] {
			for _, val := range row {
				if val > theMax {
					theMax = val
				}
			}
		}
		return
	}

	for j := 0; j < 4; j++ {
		for k := 0; k < n; k++ {
			copy(states[i][k], states[i-1][k])
		}
		push(i, j)
		simulate(i + 1)
	}

}
