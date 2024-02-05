package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &l, &c)
	charset = make([]byte, c)
	idxs = make([]int, l)
	for i := range charset {
		fmt.Fscanf(r, "%c ", &charset[i])
	}

	sort.Slice(charset, func(i, j int) bool {
		return charset[i] < charset[j]
	})

	find(w, 0)
}

var (
	l, c    int
	charset []byte
	idxs    []int

	cnts [2]int // 자, 모
)

func find(w io.Writer, i int) {
	if i == l {
		if cnts[0] < 2 || cnts[1] < 1 {
			return
		}
		for _, idx := range idxs {
			fmt.Fprintf(w, "%c", charset[idx])
		}
		fmt.Fprintln(w)
		return
	}

	j := 0
	if i > 0 {
		j = idxs[i-1] + 1
	}
	for ; j < c; j++ {
		v := charset[j]
		idx := 0
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			idx = 1
		}
		cnts[idx]++
		idxs[i] = j
		find(w, i+1)
		cnts[idx]--
	}
}
