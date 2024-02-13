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
		n, m      int
		site, pwd string
		hashmap   = make(map[string]string)
	)

	fmt.Fscan(r, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &site, &pwd)
		hashmap[site] = pwd
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &site)
		fmt.Fprintln(w, hashmap[site])
	}
}
