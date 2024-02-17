package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type trie struct {
	cnt  int
	next [26]*trie
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n    int
		s    string
		root trie
	)

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s)

		cur := &root
		last, check := 0, false
		for idx, c := range s {
			c -= 'a'
			if cur.next[c] != nil {
				cur = cur.next[c]
				last = idx
				continue
			}
			if !check {
				last = idx
				check = true
			}
			next := new(trie)
			cur.next[c] = next
			cur = next
		}
		ans := s[:last+1]
		if cur.cnt > 0 {
			ans = s + strconv.Itoa(cur.cnt+1)
		}
		fmt.Fprintln(w, ans)
		cur.cnt++
	}
}
