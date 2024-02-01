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

	var visit [1e5 + 1]bool
	var start, dst int
	fmt.Fscanf(r, "%d %d\n", &start, &dst)

	visit[start] = true
	q := [][2]int{{start, 0}}

	for {
		cur := q[0]
		q = q[1:]
		if cur[0] == dst {
			fmt.Fprintln(w, cur[1])
			break
		}

		if walkR := cur[0] + 1; walkR < len(visit) && !visit[walkR] {
			q = append(q, [2]int{walkR, cur[1] + 1})
			visit[walkR] = true
		}
		if walkL := cur[0] - 1; walkL >= 0 && !visit[walkL] {
			q = append(q, [2]int{walkL, cur[1] + 1})
			visit[walkL] = true
		}
		if tpR := cur[0] * 2; tpR < len(visit) && !visit[tpR] {
			q = append(q, [2]int{tpR, cur[1] + 1})
			visit[tpR] = true
		}
	}
}
