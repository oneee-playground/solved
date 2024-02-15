import "sort"

func solution(k int, m int, score []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(score)))
    sum := 0
    for i := 0; i < len(score)/m; i++ {
        sum += min(score[i*m:(i+1)*m]...) * m
    }
    return sum
}

func min(nums ...int) int {
    m := nums[0]
    for _, num := range nums[1:] {
        if num < m {
            m = num
        }
    }
    return m
}