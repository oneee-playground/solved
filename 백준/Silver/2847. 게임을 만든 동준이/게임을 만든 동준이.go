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
		n      int
		scores []int
	)

	fmt.Fscan(r, &n)
	scores = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &scores[i])
	}

	cnt, prev := 0, scores[len(scores)-1]
	for i := len(scores) - 2; i >= 0; i-- {
		score := scores[i]
		if score >= prev {
			cnt += score - prev + 1
			prev = prev - 1
			continue
		}
		prev = score
	}
	fmt.Fprintln(w, cnt)
}
