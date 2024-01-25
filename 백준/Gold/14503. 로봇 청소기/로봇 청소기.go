package main

import "fmt"

func main() {
	var (
		n, m, cnt       int
		locX, locY, dir int
		room            [][]int

		dx = []int{-1, 0, 1, 0}
		dy = []int{0, 1, 0, -1}
	)

	fmt.Scanf("%d %d", &n, &m)
	fmt.Scanf("%d %d %d", &locY, &locX, &dir)
	dir = len(dx) - dir - 1

	room = make([][]int, n)
	for i := range room {
		row := make([]int, m)
		for j := range row {
			fmt.Scan(&row[j])
		}
		room[i] = row
	}

	for {
		if room[locY][locX] == 0 {
			room[locY][locX] = 2
			cnt++
		}
		needCleaning := false
		for i := range dx {
			newX, newY := locX+dx[i], locY+dy[i]
			if room[newY][newX] == 0 {
				needCleaning = true
			}
		}
		if needCleaning {
			dir = (dir + 1) % len(dx)
			newX, newY := locX+dx[dir], locY+dy[dir]
			if room[newY][newX] == 0 {
				locX, locY = newX, newY
			}
			continue
		}
		newX, newY := locX-dx[dir], locY-dy[dir]
		if room[newY][newX] == 1 {
			break
		}
		locX, locY = newX, newY
	}
	fmt.Println(cnt)
}
