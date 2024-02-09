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
		bytes []byte
	)
	fmt.Fscan(r, &bytes)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] > bytes[j]
	})

	num := 0
	for _, b := range bytes {
		num += int(b)
	}

	if num%3 != 0 || bytes[len(bytes)-1] != '0' {
		fmt.Fprintln(w, -1)
		return
	}
	fmt.Fprintf(w, "%s\n", bytes)
}
