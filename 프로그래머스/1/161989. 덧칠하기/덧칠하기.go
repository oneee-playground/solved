func solution(n int, m int, section []int) int {
    if len(section) == 1 {
        return 1
    }
    
    cnt, idx, cur := 0, 1, section[0]
    for idx < len(section) {
        diff := 1
        for diff <= m && idx < len(section) {
            diff += section[idx] - cur 
            cur = section[idx]
            idx++
        }
        if diff > m {
            idx--
        }
        cnt++
    }
    return cnt
}
