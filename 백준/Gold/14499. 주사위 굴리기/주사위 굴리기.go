package main

import (
	"bufio"
	"fmt"
	"os"
)

type dice struct {
	y, x int

	up, down                 int
	north, west, south, east int
}

func (d *dice) turn(from, to *int) {
	tmp := *to
	*to = d.up
	d.up = *from
	*from = d.down
	d.down = tmp
}

func (d *dice) turnNorth() { d.turn(&d.south, &d.north) }
func (d *dice) turnSouth() { d.turn(&d.north, &d.south) }
func (d *dice) turnEast()  { d.turn(&d.west, &d.east) }
func (d *dice) turnWest()  { d.turn(&d.east, &d.west) }

var (
	N, M, K int
	board   [][]int
)

const (
	east = 1 + iota
	west
	north
	south
)

func move(d *dice, dir int) (valid bool) {
	isOut := func(y, x int) bool { return x < 0 || y < 0 || x == M || y == N }

	switch dir {
	case east:
		if isOut(d.y, d.x+1) {
			return false
		}
		d.x++
		d.turnEast()
	case west:
		if isOut(d.y, d.x-1) {
			return false
		}
		d.x--
		d.turnWest()
	case north:
		if isOut(d.y-1, d.x) {
			return false
		}
		d.y--
		d.turnNorth()
	case south:
		if isOut(d.y+1, d.x) {
			return false
		}
		d.y++
		d.turnSouth()
	}

	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	dice := dice{}

	fmt.Fscan(r, &N, &M, &dice.y, &dice.x, &K)

	board = make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(r, &board[i][j])
		}
	}

	var dir int
	for i := 0; i < K; i++ {
		fmt.Fscan(r, &dir)

		valid := move(&dice, dir)
		if !valid {
			continue
		}

		floor := board[dice.y][dice.x]
		if floor == 0 {
			board[dice.y][dice.x] = dice.down
		} else {
			dice.down = floor
			board[dice.y][dice.x] = 0
		}

		fmt.Fprintln(w, dice.up)
	}
}
