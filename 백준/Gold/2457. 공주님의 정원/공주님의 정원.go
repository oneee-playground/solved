package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type (
	flower struct {
		s, e int
	}
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n       int
		flowers []flower
	)

	fmt.Fscan(r, &n)
	flowers = make([]flower, n)
	for i := range flowers {
		var sm, sd, em, ed int
		fmt.Fscan(r,
			&sm, &sd,
			&em, &ed,
		)
		flowers[i].s = sm*100 + sd
		flowers[i].e = em*100 + ed
	}

	sort.Slice(flowers, func(i, j int) bool {
		f1, f2 := flowers[i], flowers[j]
		if f1.s != f2.s {
			return f1.s < f2.s
		}
		return f1.e > f2.e
	})

	start, end := 301, 1201
	idx, max, cnt := 0, 0, 0
	for start < end {
		found := false
		for i := idx; i < n; i++ {
			f := flowers[i]
			if f.s > start {
				break
			}
			if f.e > max {
				max = f.e
				idx = i + 1
				found = true
			}
		}
		if !found {
			break
		}
		start = max
		cnt++
	}
	if max < end {
		cnt = 0
	}
	fmt.Fprintln(w, cnt)
}
