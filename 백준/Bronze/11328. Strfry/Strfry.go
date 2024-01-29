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
		n       int
		s1, s2  string
		hashmap = make(map[rune]int)
	)
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s1, &s2)
		for _, c := range s1 {
			hashmap[c]++
		}
		for _, c := range s2 {
			hashmap[c]--
			if hashmap[c] == 0 {
				delete(hashmap, c)
			}
		}
		if len(hashmap) == 0 {
			fmt.Fprintln(w, "Possible")
			continue
		}
		fmt.Fprintln(w, "Impossible")

		clear(hashmap)
	}
}
