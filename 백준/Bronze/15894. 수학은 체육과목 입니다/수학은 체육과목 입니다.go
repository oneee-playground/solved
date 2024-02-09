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
	fmt.Fprintln(w, n*4)
}
