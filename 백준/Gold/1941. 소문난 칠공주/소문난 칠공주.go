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

	room = make([][]byte, 5)
	for i := range room {
		row := make([]byte, 5)
		fmt.Fscanf(r, "%s\n", &row)
		room[i] = row
	}
	check = make([]bool, 5*5)
	res = make([]int, 7)
	find(0)
	fmt.Fprintln(w, cnt)
}

var (
	room [][]byte

	check []bool
	res   []int
	cnt   int

	q  [][2]int
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

func checkNear() bool {
	head := q

	visit := make([]bool, len(res))
	q = append(q, [2]int{res[0] / 5, res[0] % 5})
	visit[0] = true

	cnt := 0
	for len(q) > 0 {
		p := q[0]
		cnt++
		for i := range dx {
			for j := 1; j < len(res); j++ {
				y, x := res[j]/5, res[j]%5
				if !visit[j] && p[0]+dy[i] == y && p[1]+dx[i] == x {
					q = append(q, [2]int{y, x})
					visit[j] = true
				}
			}
		}
		q = q[1:]
	}

	q = head
	return cnt == 7
}

func find(i int) {
	if i == 7 {
		cntS := 0
		for _, idx := range res {
			if room[idx/5][idx%5] == 'S' {
				cntS++
			}
		}
		if cntS < 4 || !checkNear() {
			return
		}
		cnt++
		return
	}

	j := 0
	if i > 0 {
		j = res[i-1] + 1
	}
	for ; j < len(check); j++ {
		if check[j] {
			continue
		}
		res[i] = j
		check[j] = true
		find(i + 1)
		check[j] = false
	}
}
