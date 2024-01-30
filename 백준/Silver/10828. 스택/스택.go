package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		n, num int
		s      stack
	)

	fmt.Fscanf(r, "%d\n", &n)
	for i := 0; i < n; i++ {
		line, _ := r.ReadString('\n')
		arr := strings.Split(line[:len(line)-1], " ")
		cmd := arr[0]
		switch cmd {
		case "pop":
			num = -1
			if !s.empty() {
				s, num = s.pop()

			}
			fmt.Fprintln(w, num)
		case "empty":
			num := 0
			if s.empty() {
				num = 1
			}
			fmt.Fprintln(w, num)
		case "size":
			fmt.Fprintln(w, len(s))
		case "top":
			num = -1
			if !s.empty() {
				_, num = s.pop()
			}
			fmt.Fprintln(w, num)
		case "push":
			num, _ := strconv.ParseInt(arr[1], 10, 64)
			s = s.push(int(num))
		}
	}
}
