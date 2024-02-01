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
		f, s, g, u, d int

		visit []bool
		q     [][2]int
	)
	fmt.Fscanf(r, "%d %d %d %d %d\n", &f, &s, &g, &u, &d)

	visit = make([]bool, f+1)
	q = append(q, [2]int{s, 0})
	visit[s] = true

	for len(q) > 0 {
		p := q[0]
		if p[0] == g {
			fmt.Fprintln(w, p[1])
			return
		}
		if up := p[0] + u; up <= f && !visit[up] {
			q = append(q, [2]int{up, p[1] + 1})
			visit[up] = true
		}
		if dp := p[0] - d; dp >= 1 && !visit[dp] {
			q = append(q, [2]int{dp, p[1] + 1})
			visit[dp] = true
		}
		q = q[1:]
	}
	fmt.Fprintln(w, "use the stairs")
}
