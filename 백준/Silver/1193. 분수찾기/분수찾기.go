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
		x int
	)
	fmt.Fscan(r, &x)

	if x == 1 {
		fmt.Fprintln(w, "1/1")
		return
	}

	cnt, n := 0, 0
	top, bot := 1, 1
	for x > cnt {
		n++
		cnt = n * (n + 1) / 2
	}
	num := x - (n-1)*n/2
	if n%2 == 0 {
		top = num
		bot = n - num + 1
	} else {
		top = n - num + 1
		bot = num
	}
	fmt.Fprintf(w, "%d/%d\n", top, bot)
}
