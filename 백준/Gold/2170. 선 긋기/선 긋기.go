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
		n     int
		lines [][2]int
	)

	fmt.Fscan(r, &n)
	lines = make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &lines[i][0], &lines[i][1])
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})

	sum, idx := 0, 0
	for idx < n {
		start, end := lines[idx][0], lines[idx][1]
		for i := idx + 1; i < n; i++ {
			line := lines[i]
			if line[0] > end {
				break
			}
			idx = i
			if line[1] >= end {
				end = line[1]
			}
		}
		sum += end - start
		idx++
	}
	fmt.Fprintln(w, sum)
}
