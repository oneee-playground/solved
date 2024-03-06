func solution(board []string) int {
    var (
        start, goal [2]int
        
    )
    visit := make([][]bool, len(board))
    for i, line := range board {
        for j, c := range line {
            if c == 'R' {
                start = [2]int{i, j}
            }
            if c == 'G' {
                goal = [2]int{i, j}
            }
        }
        visit[i] = make([]bool, len(line))
    }
    
    dx := [...]int{1,0,-1,0}
    dy := [...]int{0,1,0,-1}
    
    q := [][3]int{{start[0], start[1], 0}}
    for len(q) > 0 {
        pos, cnt := [2]int{q[0][0], q[0][1]}, q[0][2]
        if pos == goal {
            return cnt
        }
        for i := range dx {
            nPos := pos
            for {
                nPos[0] += dy[i]
                nPos[1] += dx[i]
                if nPos[0] < 0 || nPos[1] < 0 || nPos[0] >= len(board) || nPos[1] >= len(board[0]) {
                    break
                }
                if board[nPos[0]][nPos[1]] == 'D' {
                    break
                }
            }
            nPos[0] -= dy[i]
            nPos[1] -= dx[i]
            if visit[nPos[0]][nPos[1]] {
                continue
            }
            visit[nPos[0]][nPos[1]] = true
            q = append(q, [3]int{nPos[0], nPos[1], cnt+1})
        }
        q = q[1:]
    }
    return -1
}