package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n int

		nodes [][2]byte
	)

	fmt.Fscanln(r, &n)
	nodes = make([][2]byte, n+1)

	var p, m byte
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%c", &p)
		p -= 'A' - 1
		for j := 0; j < 2; j++ {
			fmt.Fscanf(r, " %c", &m)
			if m != '.' {
				nodes[p][j] = m - 'A' + 1
			}
		}
		fmt.Fscanln(r)
	}

	var preorder func(idx byte)
	var inorder func(idx byte)
	var postorder func(idx byte)

	print := func(c byte) { fmt.Fprintf(w, "%c", c+'A'-1) }

	preorder = func(idx byte) {
		if idx == 0 {
			return
		}
		print(idx)
		preorder(nodes[idx][0])
		preorder(nodes[idx][1])
	}
	inorder = func(idx byte) {
		if idx == 0 {
			return
		}
		inorder(nodes[idx][0])
		print(idx)
		inorder(nodes[idx][1])
	}
	postorder = func(idx byte) {
		if idx == 0 {
			return
		}
		postorder(nodes[idx][0])
		postorder(nodes[idx][1])
		print(idx)
	}

	preorder(1)
	fmt.Fprintln(w)

	inorder(1)
	fmt.Fprintln(w)

	postorder(1)
	fmt.Fprintln(w)
}
