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

	fmt.Fscan(r, &n, &m, &k)
	notebook = make([][]bool, n)
	for i := range notebook {
		notebook[i] = make([]bool, m)
	}

	for i := 0; i < k; i++ {
		fmt.Fscan(r, &row, &column)
		for i := 0; i < row; i++ {
			for j := 0; j < column; j++ {
				fmt.Fscan(r, &sticker[i][j])
			}
		}

	outer:
		for k := 0; k < 4; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if i+row > n || j+column > m {
						continue
					}
					if paste(i, j) {
						break outer
					}
				}
			}
			rotate()
		}
	}

	cnt := 0
	for i := range notebook {
		for _, ok := range notebook[i] {
			if ok {
				cnt++
			}
		}
	}
	fmt.Fprintln(w, cnt)
}

func rotate() {
	tmp := sticker
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			sticker[i][j] = tmp[row-1-j][i]
		}
	}
	row, column = column, row
}

var (
	n, m, k     int
	row, column int

	notebook [][]bool
	sticker  [10][10]bool
)

func paste(y, x int) bool {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if notebook[y+i][x+j] && sticker[i][j] {
				return false
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if sticker[i][j] {
				notebook[y+i][x+j] = true
			}
		}
	}
	return true
}
