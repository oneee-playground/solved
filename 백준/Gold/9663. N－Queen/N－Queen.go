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

	fmt.Fscan(r, &n)
	checkV = make([]bool, n)
	checkD1 = make([]bool, 2*n-1)
	checkD2 = make([]bool, 2*n-1)
	find(0)
	fmt.Fprintln(w, cnt)
}

var (
	n, cnt  int
	checkV  []bool
	checkD1 []bool
	checkD2 []bool
)

func find(i int) {
	if i == n {
		cnt++
		return
	}

	for j := 0; j < n; j++ {
		if false ||
			checkV[j] ||
			checkD1[i+j] ||
			checkD2[i-j+n-1] {
			continue
		}
		checkV[j] = true
		checkD1[i+j] = true
		checkD2[i-j+n-1] = true
		find(i + 1)
		checkV[j] = false
		checkD1[i+j] = false
		checkD2[i-j+n-1] = false
	}
}
