package main

func solution(storey int) int {

	const CRITERIA = 5
	var (
		sum  int
		move int
	)

	for storey != 0 {
		val := storey % 10
		switch {
		case val >= CRITERIA:
			if val == CRITERIA && storey/10%10 < CRITERIA {
				move = val
				storey -= move
				break
			}
			move = 10 - val
			storey += move
		default:
			move = val
			storey -= move
		}
		sum += move
		storey /= 10
	}

	return sum
}
