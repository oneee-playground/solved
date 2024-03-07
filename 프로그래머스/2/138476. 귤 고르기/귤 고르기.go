import (
    "sort"
    "container/heap"
)

type IntHeap []int
 
func (h IntHeap) Len() int {
    return len(h)
}
 
func (h IntHeap) Less(i, j int) bool {
    return h[i] > h[j]
}
 
func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
 
func (h *IntHeap) Push(element interface{}) {
    *h = append(*h, element.(int))
}
 
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    element := old[n-1]
    *h = old[0 : n-1]
    return element
}

func solution(k int, tangerine []int) int {
    h := &IntHeap{}
    
    sort.Ints(tangerine)
    
    v, cnt := -1, 0
    for _, num := range tangerine {
        if num != v {
            heap.Push(h, cnt)
            v = num
            cnt = 0
        }
        cnt++
    }
    heap.Push(h, cnt)
    
    sum, cnt := 0, 0
    for sum < k {
        cnt++
        sum += heap.Pop(h).(int)
    }
    
    return cnt
}