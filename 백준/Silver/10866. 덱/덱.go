package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var n, num int
	var q deque
	fmt.Fscanf(r, "%d\n", &n)
	for i := 0; i < n; i++ {
		tmp, _ := r.ReadString('\n')
		arr := strings.Split(tmp[:len(tmp)-1], " ")
		cmd := arr[0]
		switch cmd {
		case "pop_back":
			q, num = q.popBack()
			fmt.Fprintln(w, num)
		case "pop_front":
			q, num = q.popFront()
			fmt.Fprintln(w, num)
		case "size":
			fmt.Fprintln(w, len(q))
		case "empty":
			num := 0
			if q.empty() {
				num = 1
			}
			fmt.Fprintln(w, num)
		case "front":
			fmt.Fprintln(w, q.front())
		case "back":
			fmt.Fprintln(w, q.back())
		case "push_front":
			num, _ := strconv.ParseInt(arr[1], 10, 64)
			q = q.pushFront(int(num))
		case "push_back":
			num, _ := strconv.ParseInt(arr[1], 10, 64)
			q = q.pushBack(int(num))
		}
	}
}
