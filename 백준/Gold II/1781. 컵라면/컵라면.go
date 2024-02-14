package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type problem struct{ deadline, ramens int }

type minHeap []int

func (m *minHeap) Len() int           { return len(*m) }
func (m *minHeap) Less(i, j int) bool { return (*m)[i] < (*m)[j] }
func (m *minHeap) Swap(i, j int)      { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }
func (m *minHeap) Push(x any)         { *m = append(*m, x.(int)) }
func (m *minHeap) top() int           { return (*m)[0] }

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
		n int
		h = new(minHeap)

		problems []problem
	)

	fmt.Fscan(r, &n)
	problems = make([]problem, n)
	*h = make(minHeap, 0, n)

	for i := range problems {
		fmt.Fscan(r, &problems[i].deadline, &problems[i].ramens)
	}

	sort.Slice(problems, func(i, j int) bool {
		pi, pj := problems[i], problems[j]
		if pi.deadline == pj.deadline {
			return pi.ramens > pj.ramens
		}
		return pi.deadline < pj.deadline
	})

	for i := range problems {
		p := problems[i]
		if h.Len() < p.deadline {
			heap.Push(h, p.ramens)
			continue
		}
		if h.top() < p.ramens {
			heap.Pop(h)
			heap.Push(h, p.ramens)
		}
	}

	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	fmt.Fprintln(w, sum)
}
