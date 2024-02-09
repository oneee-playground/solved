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
		l, p, v int
	)
	for cnt := 1; ; cnt++ {
		fmt.Fscan(r, &l, &p, &v)
		if l == 0 && p == 0 && v == 0 {
			break
		}
		use := 0
		if v%p > l {
			use = v/p*l + l
		} else {
			use = v/p*l + v%p
		}
		fmt.Fprintf(w, "Case %d: %d\n", cnt, use)
	}
}
