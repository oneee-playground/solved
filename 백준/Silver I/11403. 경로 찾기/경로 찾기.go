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

		matrix [][]int
		avail  []int
		visit  []bool
	)
	fmt.Fscan(r, &n)

	matrix = make([][]int, n)
	for i := range matrix {
		row := make([]int, n)
		for j := range row {
			fmt.Fscan(r, &row[j])
		}
		matrix[i] = row
	}

	visit = make([]bool, n)
	avail = make([]int, n)

	bfs := func(idx int) {
		q := []int{idx}
		visit[idx] = true
		for len(q) > 0 {
			i := q[0]
			for j, node := range matrix[i] {
				if node == 0 {
					continue
				}
				avail[j] = 1
				if visit[j] {
					continue
				}
				visit[j] = true
				q = append(q, j)
			}
			q = q[1:]
		}
	}

	for i := range matrix {
		bfs(i)
		for _, num := range avail {
			fmt.Fprintf(w, "%d ", num)
		}
		fmt.Fprintln(w)
		clear(visit)
		clear(avail)
	}
}
