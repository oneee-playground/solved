package main

func solution(s string) int {
	var (
		targetChar                  byte
		ptr1, ptr2                  int
		matchCnt, diffCnt, totalCnt uint32
	)

	for ptr1 < len(s) {
		totalCnt++
		s = s[ptr1:]
		targetChar = s[0]

		for ptr2 = 0; ptr2 < len(s); ptr2++ {

			if s[ptr2] == targetChar {
				matchCnt++
			} else {
				diffCnt++
			}

			if matchCnt == diffCnt {
				break
			}
		}
		ptr1 = ptr2 + 1
	}
	return int(totalCnt)
}
