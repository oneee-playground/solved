package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type minHeap []int

func (m *minHeap) Len() int      { return len(*m) }
func (m *minHeap) Swap(i, j int) { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }
func (m *minHeap) Push(x any)    { *m = append(*m, x.(int)) }

func (m *minHeap) Less(i, j int) bool {
	ai, aj := abs((*m)[i]), abs((*m)[j])
	if ai == aj {
		return (*m)[i] < (*m)[j]
	}
	return ai < aj
}

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
		if v == 0 {
			v = 0
			if h.Len() > 0 {
				v = heap.Pop(h).(int)
			}
			fmt.Fprintln(w, v)
			continue
		}
		heap.Push(h, v)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
