import "sort"

func solution(cards []int) int {
    var (
        set [101]int
        groups []int
    )
    
    for idx := range set {
        set[idx] = idx
    }
    
    for _, card := range cards {
        if set[card] != card {
            continue
        }
        group := card
        cnt := 0
        for {
            cnt++
            set[card] = group
            next := cards[card-1]
            if next == group {
                break
            }
            card = next
        }
        groups = append(groups, cnt)
    }
    if len(groups) == 1 {
        return 0
    }
    sort.Sort(sort.Reverse(sort.IntSlice(groups)))
    
    return groups[0] * groups[1]
}