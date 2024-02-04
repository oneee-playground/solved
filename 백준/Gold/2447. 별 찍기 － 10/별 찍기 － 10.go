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
		n    int
		grid [][]byte
	)

	fmt.Fscan(r, &n)
	grid = make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
	}

	fill(grid, false)
	for _, row := range grid {
		fmt.Fprintf(w, "%s\n", row)
	}
}

func fill(grid [][]byte, isBlank bool) {
	if isBlank {
		for i := range grid {
			for j := range grid {
				grid[i][j] = ' '
			}
		}
		return
	}
	if len(grid) == 1 {
		grid[0][0] = '*'
		return
	}

	n := len(grid) / 3
	for i := 0; i < 3; i++ {
		section := make([][]byte, n)
		for j := 0; j < 3; j++ {
			start := j * n
			for k := range section {
				section[k] = grid[i*n+k][start : start+n]
			}
			fill(section, i == 1 && j == 1)
		}
	}
}
