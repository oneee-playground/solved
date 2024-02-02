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
		n, k int

		eta = [1e5 + 1]int{}

		q, tpq [][2]int
	)

	fmt.Fscan(r, &n, &k)
	q = append(q, [2]int{n, 1})
	eta[n] = 1

	for len(q) > 0 {
		p := q[0]

		if p[0] == k {
			fmt.Fprintln(w, p[1]-1)
			return
		}
		tpq = append(tpq, p)

		if p[0] != 0 {
			for {
				p[0] *= 2
				if p[0] > k*2 || p[0] > 1e5 {
					break
				}
				if eta[p[0]] == 0 || eta[p[0]] > p[1] {
					eta[p[0]] = p[1]
				}
				tpq = append(tpq, p)
			}
		}
		for _, p := range tpq {
			if p[0] == k {
				fmt.Fprintln(w, p[1]-1)
				return
			}
			dist := p[1] + 1
			for _, next := range [2]int{p[0] - 1, p[0] + 1} {
				if next < 0 || next > 1e5 {
					continue
				}
				if eta[next] > 0 && eta[next] <= dist {
					continue
				}
				q = append(q, [2]int{next, dist})
				eta[next] = dist
			}
		}
		tpq = tpq[:0]
		q = q[1:]
	}
}
