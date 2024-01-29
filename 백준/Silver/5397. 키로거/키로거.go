package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type stack []byte

func (s *stack) push(item byte) {
	*s = append(*s, item)
}

func (s *stack) pop() byte {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		s1, s2 stack
		s      string
		m      int
	)
	fmt.Fscan(r, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &s)
		for _, cmd := range s {
			switch cmd {
			case '<':
				if !s1.empty() {
					s2.push(s1.pop())
				}
			case '>':
				if !s2.empty() {
					s1.push(s2.pop())
				}
			case '-':
				if !s1.empty() {
					s1.pop()
				}
			default:
				s1.push(byte(cmd))
			}
		}

		for !s1.empty() {
			s2.push(s1.pop())
		}

		buf := bytes.NewBuffer(nil)
		for !s2.empty() {
			buf.WriteByte(s2.pop())
		}
		fmt.Fprintln(w, buf.String())
	}
}
