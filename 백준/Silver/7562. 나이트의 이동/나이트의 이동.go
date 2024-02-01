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
		t, l int

		dx = [8]int{1, 1, -1, -1, 2, 2, -2, -2}
		dy = [8]int{2, -2, 2, -2, 1, -1, 1, -1}

		src, dst [2]int
		q        [][3]int

		set = make(map[[2]int]struct{})
	)
	fmt.Fscanf(r, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%d\n", &l)
		fmt.Fscanf(r, "%d %d\n", &src[0], &src[1])
		fmt.Fscanf(r, "%d %d\n", &dst[0], &dst[1])

		head := q
		q = append(q, [3]int{src[0], src[1], 0})
		set[src] = struct{}{}

		for len(q) > 0 {
			p := q[0]
			if p[0] == dst[0] && p[1] == dst[1] {
				fmt.Fprintln(w, p[2])
				break
			}
			for i := range dx {
				ny, nx := p[0]+dy[i], p[1]+dx[i]
				if ny < 0 || nx < 0 || ny >= l || nx >= l {
					continue
				}
				pos := [2]int{ny, nx}
				if _, ok := set[pos]; !ok {
					q = append(q, [3]int{ny, nx, p[2] + 1})
					set[pos] = struct{}{}
				}
			}
			q = q[1:]
		}
		clear(set)
		q = head
	}
}
