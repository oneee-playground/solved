package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []int

func (s stack) push(data int) stack {
	return append(s, data)
}

func (s stack) pop() (stack, int) {
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
		k, n int
		s    stack
	)

	fmt.Fscanf(r, "%d\n", &k)
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &n)
		if n == 0 {
			s, _ = s.pop()
			continue
		}
		s = s.push(n)
	}
	sum := 0
	for !s.empty() {
		s, n = s.pop()
		sum += n
	}
	fmt.Fprintln(w, sum)
}
