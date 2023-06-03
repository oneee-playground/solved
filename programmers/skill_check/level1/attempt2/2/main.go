package main

func solution(n int, m int, section []int) int {
	var cnt, leftIdx, jdx, idx int

	for idx < len(section) {
		section = section[idx:]
		leftIdx = section[0]
		for jdx = 1; jdx < len(section); jdx++ {
			if section[jdx]-leftIdx >= m {
				break
			}
		}
		idx = jdx
		cnt++
	}

	return cnt
}
