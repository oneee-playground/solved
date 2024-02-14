package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type minHeap []int

func (m *minHeap) Len() int           { return len(*m) }
func (m *minHeap) Less(i, j int) bool { return (*m)[i] < (*m)[j] }
func (m *minHeap) Swap(i, j int)      { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }
func (m *minHeap) Push(x any)         { *m = append(*m, x.(int)) }

func (m *minHeap) Pop() any {
	last := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return last
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, v int
		h    = new(minHeap)
	)

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v)
		heap.Push(h, v)
	}

	sum := 0
	for h.Len() > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		sum += a + b
		heap.Push(h, a+b)
	}
	fmt.Fprintln(w, sum)
}
