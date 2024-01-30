package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue []int

func (q queue) push(data int) queue {
	return append(q, data)
}

func (q queue) pop() (queue, int) {
	return q[1:], q[0]
}

func (q queue) empty() bool {
	return len(q) == 0
}

func (q queue) front() int {
	return q[0]
}

func (q queue) back() int {
	return q[len(q)-1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	var q queue
	fmt.Fscanf(r, "%d\n", &n)
	for i := 0; i < n; i++ {
		tmp, _ := r.ReadString('\n')
		arr := strings.Split(tmp[:len(tmp)-1], " ")
		cmd := arr[0]
		switch cmd {
		case "pop":
			num := -1
			if !q.empty() {
				q, num = q.pop()
			}
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
			num := -1
			if !q.empty() {
				num = q.front()
			}
			fmt.Fprintln(w, num)
		case "back":
			num := -1
			if !q.empty() {
				num = q.back()
			}
			fmt.Fprintln(w, num)
		case "push":
			num, _ := strconv.ParseInt(arr[1], 10, 64)
			q = q.push(int(num))
		}
	}
}
