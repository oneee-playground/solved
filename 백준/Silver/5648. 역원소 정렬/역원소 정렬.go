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
		tmp string
	)

	fmt.Fscan(r, &n)
	arr = make([]string, n)
	for i := range arr {
		fmt.Fscan(r, &tmp)
		b := []byte(tmp)
		for i := 0; i < len(tmp)/2; i++ {
			b[i], b[len(tmp)-1-i] = b[len(tmp)-1-i], b[i]
		}
		arr[i] = string(trimLeadingZeros(b))
	}

	sort.Slice(arr, func(i, j int) bool {
		len1, len2 := len(arr[i]), len(arr[j])
		if len1 != len2 {
			return len1 < len2
		}
		return arr[i] < arr[j]
	})
	for _, s := range arr {
		fmt.Fprintln(w, s)
	}
}

func trimLeadingZeros(b []byte) []byte {
	idx := 0
	for idx < len(b)-1 && b[idx] == '0' {
		idx++
	}
	return b[idx:]
}
