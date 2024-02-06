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

	fmt.Fscan(r, &n, &m, &green, &red)
	garden = make([][]int, n)
	visit = make([][][2]int, n)
	for i := range garden {
		row := make([]int, m)
		for j := range row {
			fmt.Fscan(r, &row[j])
			if row[j] == 2 {
				available = append(available, [2]int{i, j})
			}
		}
		garden[i] = row
		visit[i] = make([][2]int, m)
	}
	ans = make([][2]int, green+red)
	find(0, green, red)
	fmt.Fprintln(w, max)
}

var (
	n, m, green, red int

	garden [][]int
	visit  [][][2]int

	available [][2]int

	ans [][2]int // idx, color(1=green, 2=red)
	max int

	q  [][4]int // y, x, color, time
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

func bfs() {
	head := q

	for _, v := range ans {
		pos := available[v[0]]
		q = append(q, [4]int{pos[0], pos[1], v[1], 0})
		visit[pos[0]][pos[1]] = [2]int{v[1], 0}
	}

	cnt := 0
	for len(q) > 0 {
		p := q[0]
		if visit[p[0]][p[1]][0] == 3 {
			q = q[1:]
			continue
		}
		for i := range dx {
			ny, nx := p[0]+dy[i], p[1]+dx[i]
			if ny < 0 || nx < 0 || ny >= n || nx >= m {
				continue
			}
			if garden[ny][nx] == 0 || visit[ny][nx][0] == 3 {
				continue
			}
			if visit[ny][nx][0] == 0 {
				q = append(q, [4]int{ny, nx, p[2], p[3] + 1})
				visit[ny][nx] = [2]int{p[2], p[3] + 1}
				continue
			}
			if p[2] != visit[ny][nx][0] && p[3]+1 == visit[ny][nx][1] {
				visit[ny][nx][0] = 3
				cnt++
			}
		}
		q = q[1:]
	}

	if cnt > max {
		max = cnt
	}

	for i := range visit {
		clear(visit[i])
	}
	q = head
}

func find(i, g, r int) {
	if i == green+red {
		bfs()
		return
	}

	j := 0
	if i > 0 {
		j = ans[i-1][0] + 1
	}
	for ; j < len(available); j++ {
		if g > 0 {
			ans[i] = [2]int{j, 1}
			find(i+1, g-1, r)
		}
		if r > 0 {
			ans[i] = [2]int{j, 2}
			find(i+1, g, r-1)
		}
	}
}
