package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	N              int
	capacity, flow [][]int
	parent         []int

	charRange = int('z'-'a'+1) * 2
)

func maxFlow(src, dst int) (result int) {
	for {
		for i := range parent {
			parent[i] = -1
		}

		queue := make([]int, 0)
		queue = append(queue, src)

		for len(queue) > 0 {
			here := queue[0]
			queue = queue[1:]
			for there := 0; there < charRange; there++ {
				if capacity[here][there]-flow[here][there] > 0 && parent[there] == -1 {
					queue = append(queue, there)
					parent[there] = here
					if there == dst {
						break
					}
				}
			}
		}

		if parent[dst] == -1 {
			break
		}

		minFlow := math.MaxInt
		for i := dst; i != src; i = parent[i] {
			minFlow = min(minFlow, capacity[parent[i]][i]-flow[parent[i]][i])
		}

		for i := dst; i != src; i = parent[i] {
			flow[parent[i]][i] += minFlow
			flow[i][parent[i]] -= minFlow
		}

		result += minFlow
	}

	return result
}

func main() {
	r := bufio.NewReader(os.Stdin)
	parent = make([]int, charRange)

	capacity = make([][]int, charRange)
	flow = make([][]int, charRange)
	for i := 0; i < charRange; i++ {
		capacity[i] = make([]int, charRange)
		flow[i] = make([]int, charRange)
	}

	fmt.Fscanln(r, &N)

	var from, to rune
	var c int
	for i := 0; i < N; i++ {
		fmt.Fscanf(r, "%c %c %d\n", &from, &to, &c)
		capacity[runeToInt(from)][runeToInt(to)] += c
		capacity[runeToInt(to)][runeToInt(from)] += c
	}

	fmt.Println(maxFlow(runeToInt('A'), runeToInt('Z')))
}

func runeToInt(c rune) int {
	if c <= 'Z' {
		return int(c - 'A')
	}
	return int(c-'a') + 26
}
