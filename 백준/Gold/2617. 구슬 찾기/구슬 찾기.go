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
		n, m    int
		in, out [][]int
	)

	fmt.Fscan(r, &n, &m)
	in = make([][]int, n+1)
	out = make([][]int, n+1)

	var a, b int
	for ; m > 0; m-- {
		fmt.Fscan(r, &a, &b)
		in[a] = append(in[a], b)
		out[b] = append(out[b], a)
	}

	q := make([]int, 0)
	visit := make([]bool, n+1)
	bfs := func(link [][]int, num int) int {
		head := q
		visit[num] = true
		q = append(q, num)

		cnt := 0
		for len(q) > 0 {
			v := q[0]
			for _, l := range link[v] {
				if visit[l] {
					continue
				}
				visit[l] = true
				q = append(q, l)
				cnt++
			}
			q = q[1:]
		}

		clear(visit)
		q = head
		return cnt
	}

	cnt, half := 0, (n-1)/2
	for i := 1; i <= n; i++ {
		big, small := bfs(out, i), bfs(in, i)
		if big > half || small > half {
			cnt++
		}
	}
	fmt.Fprintln(w, cnt)
}
