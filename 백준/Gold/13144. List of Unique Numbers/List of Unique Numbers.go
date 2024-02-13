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
		n       int
		nums    []int
		hashmap = make(map[int]struct{})
	)

	fmt.Fscan(r, &n)
	nums = make([]int, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}

	sum := 0
	left, right := 0, 0
	for right < len(nums) {
		val := nums[right]
		for {
			_, ok := hashmap[val]
			if !ok {
				break
			}
			delete(hashmap, nums[left])
			left++
		}
		hashmap[val] = struct{}{}
		sum += right - left + 1
		right++

	}
	fmt.Fprintln(w, sum)
}