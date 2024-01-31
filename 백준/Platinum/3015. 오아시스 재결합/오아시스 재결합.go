package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack [][2]int

func (s stack) push(data [2]int) stack {
	return append(s, data)
}

func (s stack) pop() (stack, [2]int) {
	last := len(s) - 1
	return s[:last], s[last]
}

func (s stack) top() [2]int {
	last := len(s) - 1
	return s[last]
}

func (s stack) empty() bool {
	return len(s) == 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		s stack

		n, num, cnt, pair int
	)
	fmt.Fscanf(r, "%d\n", &n)
	for i := 0; i < n; i++ {
		pair = 1

		fmt.Fscan(r, &num)
		for !s.empty() && s.top()[0] < num {
			cnt += s.top()[1]
			s, _ = s.pop()
		}
		if !s.empty() {
			if prev := s.top(); prev[0] == num {
				cnt += prev[1]
				if len(s) > 1 {
					cnt++
				}
				pair = s.top()[1] + 1
				s, _ = s.pop()
			} else {
				cnt++
			}
		}
		s = s.push([2]int{num, pair})
	}
	fmt.Fprintln(w, cnt)
}
