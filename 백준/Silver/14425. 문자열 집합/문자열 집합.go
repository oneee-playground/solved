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
		n, m    int
		s       string
		hashmap = make(map[string]struct{})
	)

	fmt.Fscan(r, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s)
		hashmap[s] = struct{}{}
	}

	cnt := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &s)
		_, ok := hashmap[s]
		if ok {
			cnt++
		}
	}
	fmt.Fprintln(w, cnt)
}
