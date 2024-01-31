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

	var str string
	var s stack
	for {
		str, _ = r.ReadString('\n')
		str = str[:len(str)-2]
		if len(str) == 0 {
			return
		}

		ok := true
		for _, r := range str {
			if r == '(' || r == '[' {
				s = s.push(r)
				continue
			}
			if r == ')' {
				if s.empty() || s.top() != '(' {
					ok = false
					break
				}
				s, _ = s.pop()
				continue
			}
			if r == ']' {
				if s.empty() || s.top() != '[' {
					ok = false
					break
				}
				s, _ = s.pop()
				continue
			}
		}
		if !ok || !s.empty() {
			fmt.Fprintln(w, "no")
		} else {
			fmt.Fprintln(w, "yes")
		}
		s = s[:0]
	}
}
