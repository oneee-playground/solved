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
		n   int
		arr []string
	)

	fmt.Fscan(r, &n)
	arr = make([]string, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		len1, len2 := len(arr[i]), len(arr[j])
		if len1 != len2 {
			return len1 < len2
		}
		return arr[i] < arr[j]
	})
	last := ""
	for _, s := range arr {
		if s == last {
			continue
		}
		last = s
		fmt.Fprintln(w, s)
	}
}
