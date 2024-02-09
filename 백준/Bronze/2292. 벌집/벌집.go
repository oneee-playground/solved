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
		n int
	)
	fmt.Fscan(r, &n)

	l, check := 1, 1
	for check < n {
		check += l * 6
		l++
	}
	fmt.Fprintln(w, l)
}
