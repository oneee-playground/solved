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
		t, n       int
		item, kind string
		hashmap    = make(map[string]int)
	)

	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &n)
		if n == 0 {
			fmt.Fprintln(w, 0)
			continue
		}

		for i := 0; i < n; i++ {
			fmt.Fscan(r, &item, &kind)
			hashmap[kind]++
		}

		sum := 1
		for _, num := range hashmap {
			sum *= num + 1
		}
		fmt.Fprintln(w, sum-1)

		clear(hashmap)
	}
}
