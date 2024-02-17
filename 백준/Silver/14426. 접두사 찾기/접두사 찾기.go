package main

import (
	"bufio"
	"fmt"
	"os"
)

type trie [26]*trie

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m int
		s    string
		root trie
	)

	fmt.Fscan(r, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s)
		cur := &root
		for _, c := range s {
			c -= 'a'
			if cur[c] == nil {
				cur[c] = new(trie)
			}
			cur = cur[c]
		}
	}

	cnt := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &s)

		cur := &root
		isValid := true
		for _, c := range s {
			c -= 'a'
			if cur[c] == nil {
				isValid = false
				break
			}
			cur = cur[c]
		}
		if isValid {
			cnt++
		}
	}
	fmt.Fprintln(w, cnt)
}
