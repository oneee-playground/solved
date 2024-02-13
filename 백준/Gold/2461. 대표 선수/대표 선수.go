package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type elem struct {
	num, classNum int
}

type minHeap []elem

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i].num < (*h)[j].num }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x any)         { *h = append(*h, x.(elem)) }

func (h *minHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

type class struct {
	idx   int
	stats []int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n, m    int
		classes []class
		h       = new(minHeap)
	)

	heap.Init(h)
	var mx, minDiff int = 0, 1e9

	fmt.Fscan(r, &n, &m)
	classes = make([]class, n)
	for i := range classes {
		stats := make([]int, m)
		for j := range stats {
			fmt.Fscan(r, &stats[j])
		}
		sort.Ints(stats)

		mx = max(mx, stats[0])
		heap.Push(h, elem{num: stats[0], classNum: i})
		classes[i] = class{idx: 0, stats: stats}
	}

	for {
		least := heap.Pop(h).(elem)
		minDiff = min(minDiff, mx-least.num)

		class := classes[least.classNum]
		class.idx++
		if class.idx == m {
			break
		}

		classes[least.classNum] = class

		stat := class.stats[class.idx]
		heap.Push(h, elem{num: stat, classNum: least.classNum})
		mx = max(mx, stat)
	}

	fmt.Fprintln(w, minDiff)
}
