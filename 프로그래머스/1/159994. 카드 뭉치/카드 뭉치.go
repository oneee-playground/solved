func solution(cards1 []string, cards2 []string, goal []string) string {
    var idx1, idx2 int
    for _, target := range goal {
        if idx1 < len(cards1) && cards1[idx1] == target {
            idx1++
            continue
        }
        if idx2 < len(cards2) && cards2[idx2] == target {
            idx2++
            continue
        }
        
        return "No"
    }
    
    return "Yes"
}