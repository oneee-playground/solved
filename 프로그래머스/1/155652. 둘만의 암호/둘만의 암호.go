func solution(s string, skip string, index int) string {
    ans := make([]rune, len(s))
    for idx, r := range s {
        for i := 0; i < index; {
            r++
            if r > 'z' {
                r = 'a'
            }
            if contains(skip, r) {
                continue
            }
            i++
        }
        ans[idx] = r
    }
    return string(ans)
}

func contains(s string, r rune) bool {
    for _, c := range s {
        if r == c {
            return true
        }
    }
    return false
}