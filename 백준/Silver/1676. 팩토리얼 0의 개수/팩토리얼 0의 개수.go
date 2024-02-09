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

	cnt := 0
	for i := 5; n/5 > 0; i *= 5 {
		cnt += n / 5
		n /= 5
	}
	fmt.Fprintln(w, cnt)
}
