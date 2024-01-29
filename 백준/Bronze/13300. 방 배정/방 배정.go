package main

import "fmt"

func main() {
	var n, k, s, y int
	fmt.Scan(&n, &k)
	arr := make([][2]int, 6)
	for i := 0; i < n; i++ {
		fmt.Scan(&s, &y)
		arr[y-1][s]++
	}

	cnt := 0
	for _, grade := range arr {
		for _, g := range grade {
			cnt += (g / k)
			if g % k != 0 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
