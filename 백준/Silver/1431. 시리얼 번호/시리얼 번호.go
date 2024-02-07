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
		sum1, sum2 := sumNums(arr[i]), sumNums(arr[j])
		if sum1 != sum2 {
			return sum1 < sum2
		}
		return arr[i] < arr[j]
	})
	for _, item := range arr {
		fmt.Fprintln(w, item)
	}
}

func sumNums(s string) (sum int) {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			sum += int(c - '0')
		}
	}
	return
}
