package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var input string
	r := bufio.NewReader(os.Stdin)
	input, _ = r.ReadString('\n')

	words := strings.Split(input, " ")
	var count int
	for i:= range words {
		if words[i] != "\n" && words[i] != "" {
			count++
		}
	}
	fmt.Println(count)
}