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

	fmt.Fscan(r, &n, &m)
	village = make([][]int, n)
	for i := range village {
		row := make([]int, n)
		for j := range row {
			fmt.Fscan(r, &row[j])
			if row[j] == 1 {
				houses = append(houses, [2]int{i, j})
			}
			if row[j] == 2 {
				chickens = append(chickens, [2]int{i, j})
			}
		}
		village[i] = row
	}
	idx = make([]int, m)
	pick(0)
	fmt.Fprintln(w, min)
}

var (
	n, m int

	village [][]int

	chickens [][2]int
	houses   [][2]int

	idx []int
	min int = (1 << 32) - 1
)

func eval() {
	sum := 0
	for _, house := range houses {
		tmp := 2 * n
		for _, i := range idx {
			chicken := chickens[i]
			dist := abs(house[0]-chicken[0]) + abs(house[1]-chicken[1])
			if dist < tmp {
				tmp = dist
			}
		}
		sum += tmp
	}
	if sum < min {
		min = sum
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pick(i int) {
	if i == m {
		eval()
		return
	}

	j := 0
	if i > 0 {
		j = idx[i-1] + 1
	}
	for ; j < len(chickens); j++ {
		idx[i] = j
		pick(i + 1)
	}
}
