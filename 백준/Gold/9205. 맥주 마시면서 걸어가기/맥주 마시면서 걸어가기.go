package main

import (
	"bufio"
	"fmt"
	"os"
)

var spots [][3]int // x y status

func dfs(idx int) bool {
	if idx == len(spots)-1 {
		return true
	}
	cur := spots[idx]
	spots[idx][2] = 1
	for i := 1; i < len(spots); i++ {
		spot := spots[i]
		if spot[2] == 1 {
			continue
		}
		delta := abs(cur[0]-spot[0]) + abs(cur[1]-spot[1])
		if delta > 1000 {
			continue
		}
		if ok := dfs(i); ok {
			return true
		}
		spots[idx][2] = 0
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		tc, numStore int
	)
	fmt.Fscan(r, &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscan(r, &numStore)
		spots = make([][3]int, numStore+2)
		for i := range spots {
			spot := [3]int{}
			fmt.Fscan(r, &spot[0], &spot[1])
			spots[i] = spot
		}

		happy := dfs(0)
		if happy {
			fmt.Fprintln(w, "happy")
		} else {
			fmt.Fprintln(w, "sad")
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
