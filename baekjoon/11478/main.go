package main

import (
	"fmt"
)

func main() {
	var (
		input string
		set   = make(map[string]struct{})
	)

	fmt.Scanf("%s", &input)

	for i := range input {
		for j := i; j <= len(input); j++ {
			set[input[i:j]] = struct{}{}
		}
	}

	fmt.Println(len(set) - 1)
}
