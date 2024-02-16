package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		a, b string
	)
	fmt.Fscan(r, &a, &b)
	v := 0
	if strings.Contains(a, b) {
		v = 1
	}
	fmt.Fprintln(w, v)
}
