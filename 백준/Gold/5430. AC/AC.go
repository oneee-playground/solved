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

func (q deque) back() int {
	if q.empty() {
		return -1
	}
	return q[len(q)-1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t, n int
	var p string
	fmt.Fscanf(r, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%s\n", &p)
		fmt.Fscanf(r, "%d\n", &n)
		q := make(deque, n)

		fmt.Fscanf(r, "[")
		if n > 0 {
			fmt.Fscanf(r, "%d", &q[0])
			for j := 1; j < n; j++ {
				fmt.Fscanf(r, ",%d", &q[j])
			}
		}
		fmt.Fscanf(r, "]\n")

		asc, err := true, false
		for _, cmd := range p {
			if cmd == 'R' {
				asc = !asc
				continue
			}
			if q.empty() {
				err = true
				break
			}
			if asc {
				q, _ = q.popFront()
			} else {
				q, _ = q.popBack()
			}
		}
		if err {
			fmt.Fprintln(w, "error")
			continue
		}

		fmt.Fprintf(w, "[")
		if !asc {
			for j := len(q) - 1; j >= 0; j-- {
				if j == 0 {
					fmt.Fprintf(w, "%d", q[j])
					continue
				}
				fmt.Fprintf(w, "%d,", q[j])
			}
		} else {
			for j := 0; j < len(q); j++ {
				if j == len(q)-1 {
					fmt.Fprintf(w, "%d", q[j])
					continue
				}
				fmt.Fprintf(w, "%d,", q[j])
			}
		}
		fmt.Fprintln(w, "]")
	}
}
