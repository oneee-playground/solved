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
		tmp  []rune
	)
	fmt.Fscan(r, &a, &b)

	for _, c := range a {
		if c >= '0' && c <= '9' {
			continue
		}
		tmp = append(tmp, c)
	}
	v := 0
	if strings.Contains(string(tmp), b) {
		v = 1
	}
	fmt.Fprintln(w, v)
}
