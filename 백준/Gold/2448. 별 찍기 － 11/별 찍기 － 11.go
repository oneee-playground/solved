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
		grid[i] = make([]byte, n*2)
	}

	fill(grid, n, 0, 0)
	for _, row := range grid {
		for _, c := range row {
			if c != '*' {
				c = ' '
			}
			fmt.Fprintf(w, "%c", c)
		}
		fmt.Fprintln(w)
	}
}

var item = [][]byte{
	[]byte("  *   "),
	[]byte(" * *  "),
	[]byte("***** "),
}

func fill(grid [][]byte, n, y, x int) {
	if n == 3 {
		for i := range item {
			copy(grid[y+i][x:x+6], item[i])
		}
		return
	}

	fill(grid, n/2, y, x+n/2)
	fill(grid, n/2, y+n/2, x)
	fill(grid, n/2, y+n/2, x+n)
}
