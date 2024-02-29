func solution(a int, b int, n int) int {
    sum := 0
    for n >= a {
        rem := n % a
        n = (b * (n / a))
        sum += n
        n += rem
    }
    return sum
}