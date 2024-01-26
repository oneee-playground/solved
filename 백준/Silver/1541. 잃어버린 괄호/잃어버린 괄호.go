package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		exp string
		ans int64
	)
	fmt.Scan(&exp)

	splitted := strings.Split(exp, "-")
	ans += sumAll(splitted[0])
	for _, s := range splitted[1:] {
		ans -= sumAll(s)
	}

	fmt.Println(ans)
}

func sumAll(s string) (sum int64) {
	splitted := strings.Split(s, "+")
	for _, str := range splitted {
		v, _ := strconv.ParseInt(str, 10, 64)
		sum += v
	}
	return
}
