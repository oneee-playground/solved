import "sort"

func solution(n int, lost []int, reserve []int) int {
    cnt := n - len(lost)
    
    sort.Slice(lost, func(i, j int) bool {
        return lost[i] < lost[j]    
    })
    
    sort.Slice(reserve, func(i, j int) bool {
        return reserve[i] < reserve[j]    
    })
    
    for i := 0; i < len(lost); i++ {
        for j := 0; j < len(reserve); j++ {
            if lost[i] == reserve[j] {
                lost[i] = 0
                reserve[j] = 0
                cnt++
            }
        }
    }
    
    lostIdx, reserveIdx := 0, 0
    for lostIdx != len(lost) && reserveIdx != len(reserve) {
        lostNum := lost[lostIdx]
        reserveNum := reserve[reserveIdx]
        
        if reserveNum == 0 {
            reserveIdx++
            continue
        }
        
        if lostNum == 0 {
            lostIdx++
            continue
        }
        
        if lostNum == reserveNum || lostNum == reserveNum - 1 || lostNum == reserveNum + 1 {
            cnt++
            lostIdx++
            reserveIdx++
            continue
        }
        
        if lostNum > reserveNum {
            reserveIdx++
        } else {
            lostIdx++
        }
    }
    
    return cnt
}