func solution(s string) []int {
    cnt, zeros := 0, 0
    for s != "1" {
        zero := countZero(s)
        zeros += zero
        s = createBinary(len(s) - zero)
        cnt++
    }
    
    return []int{cnt, zeros}
}

func countZero(s string) (cnt int) {
    for _, c := range s {
        if c == '0' {
            cnt++
        }
    }
    return cnt
}

func createBinary(l int) string {
    arr := make([]byte, 0)
    for l > 1 {
        if l % 2 == 0 {
            arr = append(arr, '0')
        } else {
            arr = append(arr, '1')
        }
        l /= 2
    }
    arr = append(arr, '1')
    
    // reverse
    for i := 0; i < len(arr) / 2; i++ {
        arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
    }
    
    return string(arr)
}