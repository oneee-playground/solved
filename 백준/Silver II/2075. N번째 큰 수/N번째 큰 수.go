package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type maxHeap []int

func (m *maxHeap) Len() int           { return len(*m) }
func (m *maxHeap) Less(i, j int) bool { return (*m)[i] > (*m)[j] }
func (m *maxHeap) Swap(i, j int)      { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }
func (m *maxHeap) Push(x any)         { *m = append(*m, x.(int)) }

func (m *maxHeap) Pop() any {
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
		h    = new(maxHeap)
	)
	fmt.Fscan(r, &n)
	*h = make(maxHeap, 0, n*n)
	for i := 0; i < n*n; i++ {
		fmt.Fscan(r, &v)
		h.Push(v)
	}
	heap.Init(h)
	for ; n > 0; n-- {
		v = heap.Pop(h).(int)
	}
	fmt.Fprintln(w, v)
}
