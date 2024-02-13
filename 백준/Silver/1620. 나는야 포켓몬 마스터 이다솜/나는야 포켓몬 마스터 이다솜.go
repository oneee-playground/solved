package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m    int
		s       string
		hashmap = make(map[string]int)
		names   []string
	)

	fmt.Fscan(r, &n, &m)
	names = make([]string, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &s)
		hashmap[s] = i
		names[i] = s
	}

	for ; m > 0; m-- {
		fmt.Fscan(r, &s)
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintln(w, hashmap[s])
			continue
		}
		fmt.Fprintln(w, names[num])
	}
}
