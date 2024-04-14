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

	var (
		row, col int

		pPos   [2]int
		ePoses = make(map[[2]int]int)

		grid  [][]byte
		moves string
	)

	fmt.Fscan(r, &row, &col)
	grid = make([][]byte, row)
	for i := range grid {
		row := make([]byte, col)
		fmt.Fscan(r, &row)
		for j := range row {
			switch row[j] {
			case 'I':
				pPos = [2]int{i, j}
			case 'R':
				ePoses[[2]int{i, j}] = 1
			}
			row[j] = '.'
		}
		grid[i] = row
	}

	fmt.Fscan(r, &moves)

	i := 0
	for ; i < len(moves); i++ {
		pPos = determinePNext(pPos, moves[i])
		if ePoses[pPos] == 1 {
			goto defeat
		}

		queue := make([][2]int, 0, len(ePoses))
		for pos := range ePoses {
			queue = append(queue, pos)
		}

		for _, pos := range queue {
			nPos := determineENext(pPos, pos)
			if nPos == pPos {
				goto defeat
			}

			ePoses[nPos]++
			ePoses[pos]--
		}

		for pos, cnt := range ePoses {
			if cnt != 1 {
				delete(ePoses, pos)
			}
		}
	}

	grid[pPos[0]][pPos[1]] = 'I'
	for pos := range ePoses {
		grid[pos[0]][pos[1]] = 'R'
	}

	for i := range grid {
		for j := range grid[i] {
			fmt.Fprintf(w, "%c", grid[i][j])
		}
		fmt.Fprintln(w)
	}

	return
defeat:
	fmt.Fprintf(w, "kraj %d\n", i+1)
}

var deltas = map[byte][2]int{
	'1': {1, -1},
	'2': {1, 0},
	'3': {1, 1},
	'4': {0, -1},
	'5': {0, 0},
	'6': {0, 1},
	'7': {-1, -1},
	'8': {-1, 0},
	'9': {-1, 1},
}

func determinePNext(cur [2]int, c byte) [2]int {
	delta := deltas[c]
	return [2]int{cur[0] + delta[0], cur[1] + delta[1]}
}

func determineENext(p [2]int, cur [2]int) [2]int {
	yDiff := p[0] - cur[0]
	xDiff := p[1] - cur[1]

	delta := [2]int{}
	if yDiff != 0 {
		if yDiff > 0 {
			delta[0] = 1
		} else {
			delta[0] = -1
		}
	}
	if xDiff != 0 {
		if xDiff > 0 {
			delta[1] = 1
		} else {
			delta[1] = -1
		}
	}

	return [2]int{cur[0] + delta[0], cur[1] + delta[1]}
}
