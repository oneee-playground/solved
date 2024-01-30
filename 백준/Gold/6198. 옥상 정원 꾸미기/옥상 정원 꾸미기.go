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

func (s stack) empty() bool {
	return len(s) == 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n, num int
		sum    uint64
		cur    [2]int
		s      stack
	)

	fmt.Fscanf(r, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &num)
		for !s.empty() {
			tmp, cur := s.pop()
			if cur[0] > num {
				break
			}
			sum += uint64(i - cur[1] - 1)
			s = tmp
		}
		s = s.push([2]int{num, i})
	}
	for !s.empty() {
		s, cur = s.pop()
		sum += uint64(n - cur[1] - 1)
	}
	fmt.Fprintln(w, sum)
}
