package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	dx = [4]int8{1, -1, 0, 0}
	dy = [4]int8{0, 0, 1, -1}

	n, m    int
	k       uint8
	matrix  [][]uint8
	visited [][]uint8
)

type pos struct {
	x   int
	y   int
	k   uint8
	cnt int
}

func main() {
	n, m, k, matrix = input()

	visited = make([][]uint8, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]uint8, m)
	}

	result := move()
	fmt.Println(result)
}

func move() int {
	queue := make(chan pos, n*m)
	defer close(queue)

	start := pos{0, 0, k, 1}
	visited[start.y][start.x] = start.k + 1
	queue <- start
	for p := range queue {
		if p.x == m-1 && p.y == n-1 {
			return p.cnt
		}
		for i := 0; i < 4; i++ {
			nx := p.x + int(dx[i])
			ny := p.y + int(dy[i])
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if visited[ny][nx] >= p.k+1 {
				continue
			}
			if matrix[ny][nx] == 1 {
				if p.k > 0 {
					visited[ny][nx] = p.k + 1
					queue <- pos{nx, ny, p.k - 1, p.cnt + 1}
				}
				continue
			}
			visited[ny][nx] = p.k + 1
			queue <- pos{nx, ny, p.k, p.cnt + 1}
		}
		if len(queue) == 0 {
			return -1
		}
	}
	return -1
}

func input() (int, int, uint8, [][]uint8) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	splitInput := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(splitInput[0])
	m, _ := strconv.Atoi(splitInput[1])
	k, _ := strconv.Atoi(splitInput[2])

	matrix := make([][]uint8, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		matrix[i] = make([]uint8, m)
		for index, ch := range scanner.Text() {
			num, _ := strconv.ParseUint(string(ch), 10, 8)
			matrix[i][index] = uint8(num)
		}
	}
	return n, m, uint8(k), matrix
}
