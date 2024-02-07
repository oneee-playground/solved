package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		s   string
		arr [][]byte
	)

	fmt.Fscan(r, &s)
	arr = make([][]byte, len(s))

	for i := range arr {
		arr[i] = []byte(s[i:])
	}

	slices.SortFunc(arr, func(b1, b2 []byte) int {
		return bytes.Compare(b1, b2)
	})

	for _, s := range arr {
		fmt.Fprintf(w, "%s\n", s)
	}
}
