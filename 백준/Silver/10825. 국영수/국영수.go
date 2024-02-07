package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type student struct {
	kor, eng, mat int

	name string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		n   int
		arr []student
	)

	fmt.Fscan(r, &n)
	arr = make([]student, n)

	for i := range arr {
		s := student{}
		fmt.Fscan(r, &s.name, &s.kor, &s.eng, &s.mat)
		arr[i] = s
	}

	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		switch {
		case a.kor != b.kor:
			return a.kor > b.kor
		case a.eng != b.eng:
			return a.eng < b.eng
		case a.mat != b.mat:
			return a.mat > b.mat
		}
		return a.name < b.name
	})

	for _, s := range arr {
		fmt.Fprintln(w, s.name)
	}
}
