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
		pushedLast  bool
		cnt, streak int

		str string
	)
	fmt.Fscanf(r, "%s\n", &str)
	for _, r := range str {
		if r == '(' {
			streak++
			pushedLast = true
		}
		if r == ')' {
			streak--
			if pushedLast {
				cnt += streak
			} else {
				cnt++
			}
			pushedLast = false
		}
	}
	fmt.Fprintln(w, cnt)
}
