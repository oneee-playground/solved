package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		kg  int32
		cnt int32
	)

	fmt.Fscan(bufio.NewReader(os.Stdin), &kg)

	cnt, remain := countBeforeMultipleOfFive(kg)
	if remain < 5 {
		if remain != 0 {
			cnt = -1
		}
		displayOutput(cnt)
		return
	}

	displayOutput(cnt + remain/5)
}

func countBeforeMultipleOfFive(kg int32) (cnt, remain int32) {
	for kg%5 != 0 && kg > 0 {
		cnt++
		kg -= 3
	}
	return cnt, kg
}

func displayOutput(cnt int32) {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, cnt)
	w.Flush()
}
