var moves = make([][]int, 0)

func hanoi(n, start, end, tmp int) {
    if n == 0 {
        return
    }
    hanoi(n - 1, start, tmp, end)
    move(start, end)
    hanoi(n - 1, tmp, end, start)
}

func move(start, end int) {
    moves = append(moves, []int{start, end})
}

func solution(n int) [][]int {
    hanoi(n, 1, 3, 2)
    return moves
}