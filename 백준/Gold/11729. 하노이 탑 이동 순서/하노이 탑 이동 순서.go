package main

import (
	"bufio"
	"fmt"
	"io"
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
	fmt.Fprintln(w, (1<<n)-1)
	hanoi(w, n, 1, 2, 3)
}

func hanoi(w io.Writer, n, src, tmp, dst int) {
	if n == 0 {
		return
	}
	hanoi(w, n-1, src, dst, tmp)
	fmt.Fprintln(w, src, dst)
	hanoi(w, n-1, tmp, src, dst)
}
