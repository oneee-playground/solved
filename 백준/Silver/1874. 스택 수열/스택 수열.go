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
		n, idx int
		s      stack
		moves  []byte
	)

	fmt.Fscanf(r, "%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for num := 1; num <= n; num++ {
		s = s.push(num)
		moves = append(moves, '+')

		for !s.empty() {
			tmp, peek := s.pop()
			if peek != arr[idx] {
				break
			}
			s = tmp
			moves = append(moves, '-')
			idx++
		}
	}
	if len(s) != 0 {
		fmt.Fprintln(w, "NO")
		return
	}
	for _, move := range moves {
		fmt.Fprintf(w, "%c\n", move)
	}
}
