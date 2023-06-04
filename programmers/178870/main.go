package main

func solution(sequence []int, k int) []int {
	var (
		candidates        [][2]int
		leftPtr, rightPtr int
		minIdx            int
		minLen            int = 0xFFFFFFFF / 2
		sum               int = 0
	)

	for leftPtr = range sequence {
		for rightPtr < len(sequence) {
			sum += sequence[rightPtr]
			if sum >= k {
				if sum == k {
					candidates = append(candidates, [2]int{leftPtr, rightPtr})
				}
				sum -= sequence[rightPtr]
				break
			}
			rightPtr++
		}
		sum -= sequence[leftPtr]
	}

	for i, candidate := range candidates {
		if length := getLen(candidate); length < minLen {
			minLen = length
			minIdx = i
		}
	}

	return candidates[minIdx][:]
}

func getLen(arr [2]int) int {
	return arr[1] - arr[0]
}
