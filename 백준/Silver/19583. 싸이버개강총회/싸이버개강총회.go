package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var (
		s, e, q    string
		time, name string
		students   = make(map[string]struct{})

		cnt int
	)

	sc.Scan()
	s = sc.Text()
	sc.Scan()
	e = sc.Text()
	sc.Scan()
	q = sc.Text()
	for sc.Scan() {
		time = sc.Text()
		sc.Scan()
		name = sc.Text()
		if time <= s {
			students[name] = struct{}{}
			continue
		}
		if time >= e && time <= q {
			_, ok := students[name]
			if ok {
				cnt++
				delete(students, name)
			}
		}
	}
	fmt.Fprintln(w, cnt)
}
