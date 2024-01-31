package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []rune

func (s stack) push(data rune) stack {
	return append(s, data)
}

func (s stack) pop() (stack, rune) {
	last := len(s) - 1
	return s[:last], s[last]
}

func (s stack) top() rune {
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
		t, cnt int
		str    string
		s      stack
	)
	fmt.Fscanf(r, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%s\n", &str)
		for _, r := range str {
			if !s.empty() && s.top() == r {
				s, _ = s.pop()
				continue
			}
			s = s.push(r)
		}
		if s.empty() {
			cnt++
		}
		s = s[:0]
	}
	fmt.Fprintln(w, cnt)
}
