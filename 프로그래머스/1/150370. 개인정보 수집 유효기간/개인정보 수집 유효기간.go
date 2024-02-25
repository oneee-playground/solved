import (
    "time"
    "strings"
    "strconv"
)

func solution(today string, terms []string, privacies []string) []int {
    now := parseTime(today)
    
    termMap := make(map[string]time.Time, len(terms))
    for _, term := range terms {
        strs := strings.Split(term, " ")
        char := strs[0]
        num, _ := strconv.Atoi(strs[1])
        
        termMap[char] = now.AddDate(0, -1 * num, 0)
    }
    
    res := make([]int, 0)
    for idx, p := range privacies {
        strs := strings.Split(p, " ")
        date := parseTime(strs[0])
        char := strs[1]
        if date.Before(termMap[char]) || date.Equal(termMap[char]) {
            res = append(res, idx+1)
        }
    }
    
    return res
}

func parseTime(s string) time.Time {
    t, _ := time.Parse("2006.01.02", s)
    return t
}