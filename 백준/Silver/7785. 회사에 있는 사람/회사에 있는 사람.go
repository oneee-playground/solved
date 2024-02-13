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
		n         int
		name, cmd string

		hashmap = make(map[string]struct{})
	)

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &name, &cmd)
		if cmd == "enter" {
			hashmap[name] = struct{}{}
		} else {
			delete(hashmap, name)
		}
	}

	people := make([]string, 0, len(hashmap))
	for name := range hashmap {
		people = append(people, name)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(people)))
	for _, name := range people {
		fmt.Fprintln(w, name)
	}
}
