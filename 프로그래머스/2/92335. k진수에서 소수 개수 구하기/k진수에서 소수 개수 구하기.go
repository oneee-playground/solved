import (
    "strconv"
    "math"
)

func solution(n int, k int) int {
    cnt := 0
    s := strconv.FormatInt(int64(n), k)
    start, end := 0, 0
    for start < len(s) {
        if s[start] == '0' {
            start++
            continue
        }
        for end = start; end < len(s); end++ {
            if s[end] == '0' {
                break
            }
        }
        num, _ := strconv.ParseUint(s[start:end], 10, 64)
        if isPrime(num) {
            cnt++
        }
        start = end
    }
    return cnt
}

func isPrime(n uint64) bool {
    if n < 2 {
        return false
    }
    sqrt := math.Floor(math.Sqrt(float64(n)))
    for i := uint64(2); i <= uint64(sqrt); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}