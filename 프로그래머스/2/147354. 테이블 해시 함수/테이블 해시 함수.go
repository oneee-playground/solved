import "sort"

func solution(data [][]int, col int, row_begin int, row_end int) int {
    col--; row_begin--; row_end--
    sort.Slice(data, func(i, j int) bool {
        v1, v2 := data[i][col], data[j][col]
        if v1 == v2 {
            return data[i][0] > data[j][0]
        }
        return v1 < v2
    })
    ret := 0
    for i := row_begin; i <= row_end; i++ {
        v := 0
        for _, n := range data[i] {
            v += n % (i+1)
        }
        ret ^= v
    }
    return ret
}