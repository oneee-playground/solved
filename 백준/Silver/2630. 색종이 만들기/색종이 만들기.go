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
		n     int
		paper [][]int8
	)

	fmt.Fscan(r, &n)
	paper = make([][]int8, n)
	for i := range paper {
		row := make([]int8, n)
		for j := range row {
			fmt.Fscan(r, &row[j])
		}
		paper[i] = row
	}

	for _, cnt := range calculate(paper) {
		fmt.Fprintln(w, cnt)
	}
}
func calculate(paper [][]int8) (cnts [2]int) {
	base := paper[0][0]

	for i := range paper {
		for _, v := range paper[i] {
			if base != v {
				goto shouldDiv
			}
		}
	}
	cnts[base] = 1
	return

shouldDiv:
	half := len(paper) / 2
	for i := 0; i < 2; i++ {
		section := make([][]int8, half)
		for j := 0; j < 2; j++ {
			for k := range section {
				section[k] = paper[i*half+k][j*half : (j+1)*half]
			}
			res := calculate(section)
			for i, v := range res {
				cnts[i] += v
			}
		}
	}
	return
}
