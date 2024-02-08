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
		a, b, v int
	)
	fmt.Fscan(r, &a, &b, &v)

	cnt := (v-a)/(a-b) + 1
	if (v-a)%(a-b) != 0 {
		cnt++
	}
	if a == v {
		cnt = 1
	}
	fmt.Fprintln(w, cnt)
}
