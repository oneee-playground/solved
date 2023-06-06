package main

import (
	"sort"
)

type point struct {
	x, y int
}

func solution(maps []string) []int {

	var (
		startPoint point
		results    []int
	)

	myMap := create(maps)

	for rowIdx, row := range myMap {
		for colIdx, val := range row {
			if val != 'X' {
				startPoint = point{x: colIdx, y: rowIdx}
				results = append(results, bfs(startPoint, myMap))
			}
		}
	}

	sort.Ints(results)

	if len(results) == 0 {
		results = []int{-1}
	}

	return results
}

func create(maps []string) [][]rune {
	var myMap [][]rune

	for _, str := range maps {
		myMap = append(myMap, []rune(str))
	}
	return myMap
}

func newPoint(x, y int) point {
	return point{x: x, y: y}
}

func bfs(start point, myMap [][]rune) (sum int) {
	var (
		queue []point
	)

	queue = append(queue, start)
	sum += occupy(start, myMap)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		neighbors := createNeighbors(p)

		for _, n := range neighbors {
			if n.x < len(myMap[0]) && n.y < len(myMap) &&
				n.x >= 0 && n.y >= 0 && check(n, myMap) {
				queue = append(queue, n)
				sum += occupy(n, myMap)
			}
		}
	}

	return sum
}

func createNeighbors(p point) [4]point {

	l := point{x: p.x - 1, y: p.y}
	r := point{x: p.x + 1, y: p.y}
	u := point{x: p.x, y: p.y - 1}
	d := point{x: p.x, y: p.y + 1}

	return [4]point{
		l, r, u, d,
	}
}

func occupy(p point, myMap [][]rune) int {
	val := int(myMap[p.y][p.x] - '0')
	myMap[p.y][p.x] = 'X'
	return val
}

func check(p point, myMap [][]rune) bool {
	return myMap[p.y][p.x] != 'X'
}
