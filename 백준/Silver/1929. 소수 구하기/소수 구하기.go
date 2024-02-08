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
		m, n int

		states []bool
	)
	fmt.Fscan(r, &m, &n)
	states = make([]bool, n+1)
	states[1] = true
	for i := 2; i*i <= n; i++ {
		if states[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			states[j] = true
		}
	}
	for i := m; i <= n; i++ {
		if !states[i] {
			fmt.Fprintln(w, i)
		}
	}
}
