func solution(x int, y int, n int) int {
    if x == y {
        return 0
    }
    
    hashset := make(map[int]struct{})
    
    q := [][2]int{{x, 0}}
    
    ops := []func(v int) int {
        func(v int) int {return v+n},
        func(v int) int {return v*2},
        func(v int) int {return v*3},
    }
    
    for len(q) > 0 {
        v, cnt := q[0][0], q[0][1]
        println(v, cnt)
        for _, op := range ops {
            v := op(v)
            if v == y {
                return cnt+1
            }
            
            if v < y {
                if _, ok := hashset[v]; ok {
                    continue
                }
                hashset[v] = struct{}{}
                q = append(q, [2]int{v, cnt+1})
            }
        }
        q = q[1:]
    }
    
    return -1
}