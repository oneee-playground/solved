package main

import (
	"bufio"
	"fmt"
	"os"
)

type trie struct {
	last bool
	next [10]*trie
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		t, n int
		s    string
		root trie
	)

	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &n)

		isValid := true
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &s)

			if !isValid {
				continue
			}

			cur := &root
			for _, c := range s {
				c -= '0'
				if cur.next[c] == nil {
					cur.next[c] = new(trie)
				}
				cur = cur.next[c]

				if cur.last {
					isValid = false
					break
				}
			}
			cur.last = true

			for _, v := range cur.next {
				if v != nil {
					isValid = false
					break
				}
			}
		}

		res := "NO"
		if isValid {
			res = "YES"
		}
		fmt.Fprintln(w, res)

		root = trie{}
	}
}
