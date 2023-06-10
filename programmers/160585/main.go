package main

const (
	FIRST  = 'O'
	SECOND = 'X'
)

func solution(board []string) int {

	var (
		grid [][]rune
		seq  [2]rune
	)

	seq = [2]rune{
		FIRST, SECOND,
	}

	grid = toGrid(board)

	firstCnt, secondCnt := count(seq, grid)
	if firstCnt < secondCnt || firstCnt-secondCnt > 1 {
		return 0
	}

	firstWon := checkCompletes(FIRST, grid)
	secondWon := checkCompletes(SECOND, grid)

	if firstWon && secondWon {
		return 0
	}

	if firstWon && firstCnt == secondCnt {
		return 0
	}

	if secondWon && firstCnt > secondCnt {
		return 0
	}

	return 1
}

func count(seq [2]rune, grid [][]rune) (cnt1, cnt2 int) {
	var cnts = [2]int{}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			for idx, target := range seq {
				if checkMatches(target, x, y, grid) {
					cnts[idx]++
				}
			}
		}
	}
	return cnts[0], cnts[1]
}

func toGrid(board []string) (grid [][]rune) {
	for _, line := range board {
		grid = append(grid, []rune(line))
	}
	return grid
}

func checkMatches(target rune, x, y int, grid [][]rune) bool {
	return grid[y][x] == target
}

func checkCompletes(target rune, grid [][]rune) bool {

	dirs := [][2]int{
		{1, 0},  // horizontal 1
		{-1, 0}, // horizontal 2
		{0, 1},  // vertical 1
		{0, -1}, // vertical 2
		{1, 1},  // diagonal 1
		{-1, 1}, // diagonal 2
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			for _, dir := range dirs {
				if chainMatches(target, x, y, dir[0], dir[1], grid) == 3 {
					return true
				}
			}
		}
	}

	return false
}

func chainMatches(target rune, x, y, dirX, dirY int, grid [][]rune) (cnt int) {
	if x >= 3 || x < 0 || y >= 3 || y < 0 {
		return 0
	}
	if target == grid[y][x] {
		cnt++
	}
	return cnt + chainMatches(target, x+dirX, y+dirY, dirX, dirY, grid)
}
