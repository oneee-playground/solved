func solution(clothes [][]string) int {
    hashmap := make(map[string]int)
    for _, element := range clothes {
        t := element[1]
        if _, ok := hashmap[t]; ok {
            hashmap[t]++
            continue
        }
        hashmap[t] = 1
    }
    
    sum := 1
    for _, amount := range hashmap {
        sum *= amount + 1
    }
    sum--
    
    return sum
}