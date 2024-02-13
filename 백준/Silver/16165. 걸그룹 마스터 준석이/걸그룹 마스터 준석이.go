package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m, cnt      int
		group, name, q string
		members        = make(map[string][]string)
		groups         = make(map[string]string)
	)

	fmt.Fscan(r, &n, &m)

	groups = make(map[string]string)
	members = make(map[string][]string)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &group, &cnt)
		for j := 0; j < cnt; j++ {
			fmt.Fscan(r, &name)
			members[group] = append(members[group], name)
			groups[name] = group
		}
		sort.Strings(members[group])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(r, &q, &cnt)
		if cnt == 1 {
			fmt.Fprintln(w, groups[q])
			continue
		}
		for _, name := range members[q] {
			fmt.Fprintln(w, name)

		}
	}
}
