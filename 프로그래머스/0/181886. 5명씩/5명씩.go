func solution(names []string) []string {
    ans := make([]string, 0, len(names) / 5)
    for idx := 0; idx < len(names); idx += 5 {
        ans = append(ans, names[idx])
    }
    return ans
}