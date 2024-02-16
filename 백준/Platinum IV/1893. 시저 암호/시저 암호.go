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
		t int
	)

	fmt.Fscanln(r, &t)

	readline := func() string {
		v, _ := r.ReadString('\n')
		return v[:len(v)-1]
	}

	for ; t > 0; t-- {
		charset := readline()
		target := []byte(readline())
		src := readline()

		f := prepare(target)
		charsetHash := make(map[byte]int, len(charset))
		for idx, char := range charset {
			charsetHash[byte(char)] = idx
		}

		ans := make([]int, 0)
		for round := 0; round < len(charset); round++ {
			contains := false

			j := 0
			for i := 0; i < len(src); i++ {
				for j > 0 && src[i] != target[j] {
					j = f[j-1]
				}
				if src[i] == target[j] {
					j++
				}
				if j == len(target) {
					if contains {
						contains = false
						break
					}
					contains = true
					j = f[j-1]
				}
			}

			if contains {
				ans = append(ans, round)
			}

			for i := 0; i < len(target); i++ {
				shifted := charsetHash[target[i]] + 1
				target[i] = charset[shifted%len(charset)]
			}
		}

		switch len(ans) {
		case 0:
			fmt.Fprintln(w, "no solution")
		case 1:
			fmt.Fprintln(w, "unique:", ans[0])
		default:
			fmt.Fprint(w, "ambiguous:")
			for _, shift := range ans {
				fmt.Fprintf(w, " %d", shift)
			}
			fmt.Fprintln(w)
		}

	}

}

func prepare(s []byte) []int {
	f := make([]int, len(s))
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = f[j-1]
		}
		if s[i] == s[j] {
			j++
			f[i] = j
		}
	}
	return f
}
