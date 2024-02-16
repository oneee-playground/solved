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
		src, _    = r.ReadString('\n')
		target, _ = r.ReadString('\n')
	)

	src = src[:len(src)-1]
	target = target[:len(target)-1]
	f := prepare(target)

	cnt, j := 0, 0
	res := make([]int, 0)
	for i := 0; i < len(src); i++ {
		for j > 0 && src[i] != target[j] {
			j = f[j-1]
		}
		if src[i] == target[j] {
			j++
		}
		if j == len(target) {
			cnt++
			res = append(res, i-len(target)+1)
			j = f[j-1]
		}
	}
	fmt.Fprintln(w, cnt)
	for _, num := range res {
		fmt.Fprintf(w, "%d ", num+1)
	}
}

func prepare(s string) []int {
	f := make([]int, len(s))
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = f[j-1]
		}
		if s[i] == s[j] {
			j++
			f[i] = j
		}
	}
	return f
}
