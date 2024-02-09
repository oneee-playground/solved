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
		t, n int
	)
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &n)
		fmt.Fprintln(w, int(math.Floor(math.Sqrt(float64(n)))))
	}
}
