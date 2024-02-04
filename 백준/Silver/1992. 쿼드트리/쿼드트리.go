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
		n     int
		video [][]byte
	)

	fmt.Fscan(r, &n)
	video = make([][]byte, n)
	for i := range video {
		row := make([]byte, n)
		fmt.Fscan(r, &row)
		video[i] = row
	}
	fmt.Fprintf(w, "%s\n", compress(video))
}

func compress(video [][]byte) (res []byte) {
	base := video[0][0]
	for i := range video {
		for _, b := range video[i] {
			if b != base {
				goto div
			}
		}
	}
	return []byte{base}

div:
	res = []byte{'('}
	defer func() { res = append(res, ')') }()

	half := len(video) / 2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			start := j * half
			section := make([][]byte, half)
			for k := range section {
				section[k] = video[i*half+k][start : start+half]
			}

			res = append(res, compress(section)...)
		}
	}

	return
}
