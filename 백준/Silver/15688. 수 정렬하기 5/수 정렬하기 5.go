package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n   int
		arr []int
	)

	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	arr = make([]int, n)
	for i := range arr {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		arr[i] = n
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})
	for _, n := range arr {
		w.WriteString(strconv.FormatInt(int64(n), 10) + "\n")
	}
}
