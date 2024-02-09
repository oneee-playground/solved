package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		t, x, y int
	)

	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &x, &y)
		dist := y - x

		max := math.Trunc(math.Sqrt(float64(dist)))
		maxInt := int(max)

		res := 0
		switch {
		case max == math.Sqrt(float64(dist)):
			res = maxInt*2 - 1
		case dist <= (maxInt*maxInt)+maxInt:
			res = maxInt * 2
		default:
			res = maxInt*2 + 1
		}
		fmt.Fprintln(w, res)
	}
}
