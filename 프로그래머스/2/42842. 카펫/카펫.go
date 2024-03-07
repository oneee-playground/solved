func solution(brown int, yellow int) []int {
    for row := 1; row * row <= yellow; row++ {
        if yellow % row != 0 {
            continue
        }
        col := yellow / row
        
        expected := row * 2 + col * 2 + 4
        if brown == expected {
            return []int{col+2, row+2}
        }
    }
    return []int{}
}