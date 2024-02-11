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
		n int
	)

	fmt.Fscan(r, &n)
	if n == 1 {
		fmt.Fprintln(w, 0)
		return
	}
	primes := makePrimes(n)

	cnt := 0

	right, sum := 0, primes[0]
loop:
	for left := 0; left < len(primes); left++ {
		for right < len(primes) && sum < n {
			right++
			if right == len(primes) {
				break loop
			}
			sum += primes[right]
		}
		if sum == n {
			cnt++
		}
		sum -= primes[left]
	}
	fmt.Fprintln(w, cnt)
}

func makePrimes(n int) []int {
	nums := make([]bool, n+1)
	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if nums[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * 2; j <= n; j += i {
			nums[j] = true
		}
	}
	return primes
}
