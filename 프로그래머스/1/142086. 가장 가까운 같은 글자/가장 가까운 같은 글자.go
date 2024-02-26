func solution(s string) []int {
    res := make([]int, 0, len(s))
    
    hashmap := make(map[rune]int)
    for idx, r := range s {
        dist := -1
        
        i, ok := hashmap[r]
        if ok {
            dist = idx - i
        }
        
        hashmap[r] = idx
        res = append(res, dist)
    } 
    
    return res
}