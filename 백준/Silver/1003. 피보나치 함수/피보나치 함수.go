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
		t, n int
	)
	fmt.Fscan(r, &t)
	fibs := [41][2]int{
		{1, 0},
		{0, 1},
	}
	var fib func(num int) [2]int
	fib = func(num int) [2]int {
		v := fibs[num]
		if v != [2]int{} {
			return v
		}
		a := fib(num - 1)
		b := fib(num - 2)
		fibs[num] = [2]int{a[0] + b[0], a[1] + b[1]}
		return fibs[num]
	}

	for ; t > 0; t-- {
		fmt.Fscan(r, &n)
		res := fib(n)
		fmt.Fprintln(w, res[0], res[1])
	}
}
