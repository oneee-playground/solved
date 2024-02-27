var stat = [4][2]byte{
    {'R', 'T'},
    {'C', 'F'},
    {'J', 'M'},
    {'A', 'N'},
}

func solution(survey []string, choices []int) string {
    var (
        hashmap = make(map[byte]int, 8)
    )
    
    for idx, opt := range survey {
        first, second := opt[0], opt[1]
        bias := choices[idx] - 4
        if bias < 0 {
            hashmap[second] -= bias
            continue
        }
        hashmap[first] += bias
    }
    
    var s []byte
    for _, b := range stat {
        first, second := b[0], b[1]
        if hashmap[first] <= hashmap[second] {
            s = append(s, first)
            continue
        } 
        s = append(s, second)
    }
    return string(s)
}