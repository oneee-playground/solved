package main

func solution(lottos []int, win_nums []int) []int {

	var matchCnt, candidateCnt int32

	for _, num := range lottos {
		if num == 0 {
			candidateCnt++
			continue
		}
		if contains(win_nums, num) {
			matchCnt++
		}
	}

	return []int{
		createRank(matchCnt + candidateCnt),
		createRank(matchCnt),
	}
}

func createRank(matchCnt int32) int {
	if matchCnt < 2 {
		return 6
	}
	return 7 - int(matchCnt)
}

func contains(slice []int, target int) bool {
	for _, data := range slice {
		if data == target {
			return true
		}
	}
	return false
}
