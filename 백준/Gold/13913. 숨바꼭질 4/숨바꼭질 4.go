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

		q [][2]int
	)
	head := q

	fmt.Fscan(r, &n, &k)
	q = append(q, [2]int{n, 1})
	eta[n] = 1

	for len(q) > 0 {
		p := q[0]
		if p[0] == k {
			break
		}

		dist := p[1] + 1
		for _, next := range [3]int{p[0] - 1, p[0] + 1, p[0] * 2} {
			if next < 0 || next > 1e5 {
				continue
			}
			if eta[next] > 0 {
				continue
			}
			q = append(q, [2]int{next, dist})
			eta[next] = dist
		}
		q = q[1:]
	}
	q = head

	fmt.Fprintln(w, eta[k]-1)

	q = append(q, [2]int{k, eta[k]})
	for i := 0; q[i][1] != 1; i++ {
		p := q[i]
		if left := p[0] - 1; left >= 0 && eta[left] == p[1]-1 {
			q = append(q, [2]int{left, eta[left]})
			continue
		}
		if right := p[0] + 1; right <= 1e5 && eta[right] == p[1]-1 {
			q = append(q, [2]int{right, eta[right]})
			continue
		}
		if div := p[0] / 2; p[0]%2 == 0 && eta[div] == p[1]-1 {
			q = append(q, [2]int{div, eta[div]})
			continue
		}
		panic("wtf")
	}

	for i := len(q) - 1; i >= 0; i-- {
		fmt.Fprintf(w, "%d ", q[i][0])
	}
}
