package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var (
		n int
		meetings [][2]int
	)

	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	meetings = make([][2]int, n)
	for i := range meetings {
		meeting := [2]int{}
		fmt.Fscan(r, &meeting[0], &meeting[1])
		meetings[i] = meeting
	}

	sort.Slice(meetings, func(i, j int) bool {
		end1, end2 :=  meetings[i][1], meetings[j][1]
		if end1 == end2 {
			return meetings[i][0] < meetings[j][0]
		}
		return end1 < end2
	})

	cnt := 0
	lastEnd := 0
	for _, meeting := range meetings {
		if meeting[0] >= lastEnd {
			cnt++
			lastEnd = meeting[1]
		}
	}
	fmt.Println(cnt)
}
