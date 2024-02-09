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
		t int

		m, n, x, y int
	)
	fmt.Fscan(r, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(r, &m, &n, &x, &y)

		x = zeroIfEqual(x, m)
		y = zeroIfEqual(y, n)

		lim := lcm(m, n)

		year := -1
		for i := x; i <= lim; i += m {
			if i == 0 {
				continue
			}
			if i%n == y {
				year = i
				break
			}
		}
		fmt.Fprintln(w, year)
	}
}

func zeroIfEqual(x, y int) int {
	if x == y {
		return 0
	}
	return x
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
