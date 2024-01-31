package main

import (
	"bufio"
	"fmt"
	"os"
)

type deque [][2]int

func (q deque) pushFront(data [2]int) deque {
	return append([][2]int{data}, q...)
}

func (q deque) pushBack(data [2]int) deque {
	return append(q, data)
}

func (q deque) popFront() (deque, [2]int) {
	return q[1:], q[0]
}

func (q deque) popBack() (deque, [2]int) {
	last := len(q) - 1
	return q[:last], q[last]
}

func (q deque) empty() bool {
	return len(q) == 0
}

func (q deque) front() [2]int {
	return q[0]
}

func (q deque) back() [2]int {
	return q[len(q)-1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, l, num int
	var d deque
	fmt.Fscanf(r, "%d %d\n", &n, &l)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &num)
		for !d.empty() && d.back()[0] > num {
			d, _ = d.popBack()
		}
		d = d.pushBack([2]int{num, i})
		if d.front()[1] < i-(l-1) {
			d, _ = d.popFront()
		}
		fmt.Fprintf(w, "%d ", d.front()[0])
	}
}
