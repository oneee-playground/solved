package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	idx int
	h   uint64
}

type stack []pair

func (s stack) push(data pair) stack {
	return append(s, data)
}

func (s stack) top() pair {
	last := len(s) - 1
	return s[last]
}

func (s stack) pop() (stack, pair) {
	last := len(s) - 1
	return s[:last], s[last]
}

func (s stack) empty() bool {
	return len(s) == 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n int

		idx  int
		s    stack
		item pair

		num, maxArea, width uint64
	)

	for {
		fmt.Fscan(r, &n)
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &num)
			idx = i

			for !s.empty() && s.top().h > num {
				s, item = s.pop()
				width = uint64(i)
				if !s.empty() {
					width -= uint64(item.idx)
				}
				maxArea = max(maxArea, item.h*width)
				idx = item.idx
			}
			s = s.push(pair{idx: idx, h: num})
		}

		for !s.empty() {
			s, item = s.pop()
			width = uint64(n)
			if !s.empty() {
				width -= uint64(item.idx)
			}
			maxArea = max(maxArea, item.h*width)
		}
		fmt.Fprintln(w, maxArea)

		s = s[:0]
		maxArea = 0
	}
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
