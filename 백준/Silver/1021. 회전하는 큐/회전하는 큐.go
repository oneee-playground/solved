package main

import (
	"bufio"
	"fmt"
	"os"
)

type deque []int

func (q deque) pushFront(data int) deque {
	return append([]int{data}, q...)
}

func (q deque) pushBack(data int) deque {
	return append(q, data)
}

func (q deque) popFront() (deque, int) {
	if q.empty() {
		return q, -1
	}
	return q[1:], q[0]
}

func (q deque) popBack() (deque, int) {
	if q.empty() {
		return q, -1
	}
	last := len(q) - 1
	return q[:last], q[last]
}

func (q deque) empty() bool {
	return len(q) == 0
}

func (q deque) front() int {
	if q.empty() {
		return -1
	}
	return q[0]
}

// func (q deque) back() int {
// 	if q.empty() {
// 		return -1
// 	}
// 	return q[len(q)-1]
// }

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m, idx, num int
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	q := make(deque, n)
	for i := 0; i < n; i++ {
		q[i] = i
	}

	cnt := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &idx)
		idx--
		realIdx := 0
		for j := 0; j < len(q); j++ {
			if q[j] == idx {
				realIdx = j
				break
			}
		}
		for idx != q.front() {
			right := realIdx <= len(q)/2
			if right {
				q, num = q.popFront()
				q = q.pushBack(num)
			} else {
				q, num = q.popBack()
				q = q.pushFront(num)
			}
			cnt++
		}
		q, _ = q.popFront()
	}
	fmt.Fprintln(w, cnt)
}
