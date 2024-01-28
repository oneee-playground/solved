func solution(s string) bool {
    if len(s) != 4 && len(s) != 6 {
        return false
    }
    for _, c := range s {
        if c < '0' || c > '9' {
            return false
        } 
    }
    return true
}