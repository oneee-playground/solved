func solution(n int) int64 {

	var a, b int64 = 0, 1
	for i := 0; i <= n; i++ {
		a, b = b, (a+b)%1234567
	}

	return a
}
