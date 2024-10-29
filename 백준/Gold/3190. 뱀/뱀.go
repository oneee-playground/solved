package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	right = iota
	down
	left
	up
)

const (
	wall = 4 << iota
	apple
	seen
)

func turnLeft(cur int) int  { return (cur - 1) & 3 }
func turnRight(cur int) int { return (cur + 1) & 3 }

var dirs = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var (
	N, K, L int
	board   [][]int
	turns   []turnStruct
)

type turnStruct struct {
	x int
	d rune
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &N)
	board = make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}

	fmt.Fscan(r, &K)
	var row, col int
	for i := 0; i < K; i++ {
		fmt.Fscan(r, &row, &col)
		board[row-1][col-1] = apple
	}

	fmt.Fscan(r, &L)
	turns = make([]turnStruct, L)
	for i := 0; i < L; i++ {
		tmp := turnStruct{}
		fmt.Fscanf(r, "\n%d %c", &tmp.x, &tmp.d)
		turns[i] = tmp
	}

	head, tail := [2]int{0, 0}, [2]int{0, 0}
	dir := right

	curTurnAt := 0
	curTurn := turnStruct{-1, 0}
	if len(turns) > 0 {
		curTurn = turns[0]
	}

	var i int
	for {
		if i == curTurn.x {
			switch curTurn.d {
			case 'L':
				dir = turnLeft(dir)
			case 'D':
				dir = turnRight(dir)
			}
			curTurnAt++
			if curTurnAt < len(turns) {
				curTurn = turns[curTurnAt]
			}
		}

		board[head[0]][head[1]] = dir | wall

		delta := dirs[dir]
		head[0] += delta[0]
		head[1] += delta[1]

		if head[0] < 0 || head[1] < 0 || head[0] == N || head[1] == N {
			break
		}

		dst := board[head[0]][head[1]]
		if dst&wall != 0 {
			break
		}

		if dst&apple == 0 {
			src := board[tail[0]][tail[1]]
			board[tail[0]][tail[1]] = seen

			delta := dirs[src&3]
			tail[0] += delta[0]
			tail[1] += delta[1]
		}

		i++
	}

	fmt.Println(i + 1)
}
