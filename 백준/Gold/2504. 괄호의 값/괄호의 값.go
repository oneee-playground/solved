package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		str string
		sum int
	)
	fmt.Fscanf(r, "%s\n", &str)
	for idx := 0; idx < len(str); {
		i, num := calculate(str[idx:])
		if num == 0 {
			sum = 0
			break
		}
		idx += i
		sum += num
	}

	fmt.Fprintln(w, sum)
}
func calculate(str string) (idx, val int) {
	start := str[0]
	end := byte(0)
	switch start {
	case '(':
		end = ')'
		val = 2
	case '[':
		end = ']'
		val = 3
	}
	if len(str) < 2 {
		return 0, 0
	}
	if str[1] == end {
		return 2, val
	}

	sum := 0
	for idx = 1; idx < len(str); {
		r := str[idx]
		if r == end {
			return idx + 1, sum * val
		}
		if r == '(' || r == '[' {
			i, num := calculate(str[idx:])
			if num == 0 {
				return 0, 0
			}
			idx += i
			sum += num
			continue
		}
		break
	}
	return 0, 0
}
