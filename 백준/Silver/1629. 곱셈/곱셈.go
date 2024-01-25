package main

import (
	"fmt"
)

var (
    a, c uint64
)

func pow(v int) uint64 {
    if v == 0 {
        return 1
    }
    tmp := pow(v/2)
    tmp = tmp * tmp % c
    if (v % 2 == 0) {
        return tmp
    }
    return a * tmp % c
}

func main() {
	var b   int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println(pow(b))
}
