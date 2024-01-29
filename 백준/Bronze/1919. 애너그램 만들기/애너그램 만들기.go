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
		s1, s2  string
		hashmap = make(map[rune]int)
	)
	fmt.Fscan(r, &s1, &s2)
	for _, c := range s1 {
		hashmap[c]++
	}
	for _, c := range s2 {
		hashmap[c]--
	}
	cnt := 0
	for _, n := range hashmap {
		if n < 0 {
			n = -n
		}
		cnt += n
	}
	fmt.Println(cnt)
}
