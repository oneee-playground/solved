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
		t, n int

		line  []int
		visit []bool
		s     []int
	)

	dfs := func(i int) int {
		head := s
		defer func() { s = head }()
		s = append(s, i)

		for {
			l := len(s) - 1
			cur := s[l]
			visit[cur] = true

			next := line[cur] - 1
			if visit[next] {
				if next == i {
					return len(s)
				}
				i = next
				break
			}
			s = append(s, next)
		}

		for j := 0; j < len(s); j++ {
			if s[j] == i {
				return len(s) - j
			}
		}
		return 0
	}

	fmt.Fscan(r, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(r, &n)
		line = make([]int, n)
		visit = make([]bool, n)
		for i := range line {
			fmt.Fscan(r, &line[i])
		}

		sum := 0
		for i := range line {
			if !visit[i] {
				sum += dfs(i)
			}
		}
		fmt.Fprintln(w, n-sum)
	}
}
