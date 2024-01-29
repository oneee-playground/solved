package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	
	arr := make([]int, 'z' - 'a' + 1)
	for _, c := range s {
		arr[c - 'a']++
	}

	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
}
