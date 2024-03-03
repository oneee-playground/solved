import "sort"

func solution(arr []int, divisor int) []int {
    res := make([]int, 0)
    for _, v := range arr {
        if v % divisor == 0 {
            res = append(res, v)
        }
    }
    if len(res) == 0 {
        return []int{-1}
    }
    sort.Ints(res)
    return res
}