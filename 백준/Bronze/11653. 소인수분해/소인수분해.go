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
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			fmt.Fprintln(w, i)
			n /= i
		}
	}
	if n != 1 {
		fmt.Fprintln(w, n)
	}
}
