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

	fmt.Fscan(r, &n, &p, &q)
	table[0] = 1

	fmt.Fprintln(w, dp(n))
}

var (
	n, p, q uint64
	table   = make(map[uint64]uint64)
)

func dp(n uint64) uint64 {
	if v, ok := table[n]; ok {
		return v
	}

	val := dp(n/p) + dp(n/q)
	table[n] = val
	return val
}
