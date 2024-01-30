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
		s      stack
	)

	fmt.Fscanf(r, "%d\n", &n)
	received := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &num)
		for {
			if s.empty() {
				received[i] = 0
				break
			}
			tmp, cur := s.pop()
			if cur[0] >= num {
				received[i] = cur[1]
				break
			}
			s = tmp
		}
		s = s.push([2]int{num, i + 1})
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d ", received[i])
	}

}
