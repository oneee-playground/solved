package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type theHeap []int

func (m *theHeap) Len() int      { return len(*m) }
func (m *theHeap) Swap(i, j int) { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }
func (m *theHeap) Push(x any)    { *m = append(*m, x.(int)) }
func (m *theHeap) top() int      { return (*m)[0] }

func (m *theHeap) Pop() any {
	last := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return last
}

type (
	minHeap struct{ theHeap }
	maxHeap struct{ theHeap }
)

func (m *minHeap) Less(i, j int) bool { return m.theHeap[i] < m.theHeap[j] }
func (m *maxHeap) Less(i, j int) bool { return m.theHeap[i] > m.theHeap[j] }

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, v       int
		minh, maxh = new(minHeap), new(maxHeap)
	)

	fmt.Fscan(r, &n)
	minh.theHeap = make([]int, 0, n)
	maxh.theHeap = make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v)

		if maxh.Len() == 0 || maxh.Len() == minh.Len() {
			heap.Push(maxh, v)
		} else {
			heap.Push(minh, v)
		}

		if maxh.Len() != 0 && minh.Len() != 0 && !(maxh.top() <= minh.top()) {
			a, b := heap.Pop(maxh).(int), heap.Pop(minh).(int)
			heap.Push(maxh, b)
			heap.Push(minh, a)
		}

		fmt.Fprintln(w, maxh.top())
	}
}
