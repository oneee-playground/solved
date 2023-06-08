package main

func solution(order []int) int {
	var (
		boxNum, cnt int
		lenEdges    int
		edges       []int
	)

	boxNum = order[0]

	for _, idx := range order {
		if idx > boxNum {
			edges = append(edges, boxNum-1)
			lenEdges++
			boxNum = idx
		}

		if idx == boxNum-1 {
			boxNum--
		}

		for i := lenEdges - 1; i >= 0; i-- {
			data := edges[i]
			if data == idx {
				boxNum = data
				edges = edges[:lenEdges-1]
				lenEdges--
				break
			}
		}

		if idx != boxNum {
			break
		}
		cnt++
	}
	return cnt
}
