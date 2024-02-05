package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscanf(r, "%d\n", &n)
	eggs = make([][2]int, n)
	for i := range eggs {
		fmt.Fscanf(r, "%d %d\n", &eggs[i][0], &eggs[i][1])
	}
	find(w, 0)
	fmt.Fprintln(w, max)
}

var (
	n int

	eggs [][2]int // 내구도, 무게
	max  int
)

func find(w io.Writer, i int) {
	if i == n {
		brokeCnt := 0
		for _, egg := range eggs {
			if egg[0] <= 0 {
				brokeCnt++
			}
		}
		if brokeCnt > max {
			max = brokeCnt
		}
		return
	}
	if eggs[i][0] <= 0 {
		find(w, i+1)
		return
	}

	found := false
	for j := 0; j < n; j++ {
		if j == i || eggs[j][0] <= 0 {
			continue
		}
		eggs[j][0] -= eggs[i][1]
		eggs[i][0] -= eggs[j][1]
		find(w, i+1)
		eggs[j][0] += eggs[i][1]
		eggs[i][0] += eggs[j][1]
		found = true
	}
	if !found {
		find(w, n)
	}
}
