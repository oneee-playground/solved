func solution(wallpaper []string) []int {
    var (
        top, bottom = -1, -1
        left = 50
        right = -1
    )
    
    for i, line := range wallpaper {
        for j, c := range line {
            if c == '#' {
                if top == -1 {
                    top = i
                }
                bottom = i
                if j > right {
                    right = j
                }
                if j < left {
                    left = j
                }
            }
        }
    }
    return []int{top, left, bottom+1, right+1}
}