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
		n, y, x int
	)

	fmt.Fscan(r, &n, &y, &x)
	fmt.Fprintln(w, z(n, y, x))
}

func z(n, y, x int) int {
	if n == 0 {
		return 0
	}
	half := 1 << (n - 1)
	if y < half && x < half {
		return z(n-1, y, x)
	}
	if y < half && x >= half {
		return half*half + z(n-1, y, x-half)
	}
	if y >= half && x < half {
		return 2*half*half + z(n-1, y-half, x)
	}
	return 3*half*half + z(n-1, y-half, x-half)
}
