package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type trie struct {
	s    string
	next []*trie
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m int
		s    string
		root = trie{next: make([]*trie, 0)}
	)

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &m)

		cur := &root
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &s)
			var next *trie
			for _, n := range cur.next {
				if n.s == s {
					next = n
				}
			}
			if next == nil {
				next = &trie{s: s, next: make([]*trie, 0)}
				cur.next = append(cur.next, next)
			}
			cur = next
		}
	}

	var dfs func(t []*trie, depth int)
	dfs = func(t []*trie, depth int) {
		if len(t) == 0 {
			return
		}

		sort.Slice(t, func(i, j int) bool {
			return t[i].s < t[j].s
		})

		prefix := strings.Repeat("--", depth)

		for _, tr := range t {
			fmt.Fprintf(w, "%s%s\n", prefix, tr.s)
			dfs(tr.next, depth+1)
		}
	}

	dfs(root.next, 0)
}
