func solution(maps []string) int {
    grid = maps
    for i, line := range grid {
        for j, c := range line {
            if c == 'S' {
                sY, sX = i, j
            }
            if c == 'L' {
                lY, lX = i, j
            }
        }
    }
    
    toLever := bfs(sY, sX, 'L')
    if toLever == -1 {
        return -1
    }
    
    toExit := bfs(lY, lX, 'E')
    if toExit == -1 {
        return -1
    }
    
    return toLever + toExit
}

var (
    dx = [4]int{1, 0, -1, 0}
    dy = [4]int{0, 1, 0, -1}
    
    sX, sY int
    lX, lY int
    
    grid  []string
)

func bfs(y, x int, target byte) int {
    visit := make([][]bool, len(grid))
    for i := range visit {
        visit[i] = make([]bool, len(grid[i]))
    }
    
    q := [][3]int{{y, x, 0}}
    for len(q) > 0 {
        v := q[0]
        for i := range dx {
            ny, nx := v[0] + dy[i], v[1] + dx[i]
            if ny < 0 || nx < 0 || ny >= len(grid) || nx >= len(grid[0]) {
                continue
            }
            if visit[ny][nx] || grid[ny][nx] == 'X' {
                continue
            }
            if grid[ny][nx] == target {
                return v[2] + 1
            }
            visit[ny][nx] = true
            q = append(q, [3]int{ny, nx, v[2]+1})
        }
        q = q[1:]
    }
    return -1
}