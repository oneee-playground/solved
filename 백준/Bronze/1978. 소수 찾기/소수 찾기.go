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
		n   int
		cnt int
	)
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(r, &num)
		if isPrime(num) {
			cnt++
		}
	}
	fmt.Fprintln(w, cnt)
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
