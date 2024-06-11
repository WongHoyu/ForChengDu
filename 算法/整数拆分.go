package 算法

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}

	res := 1
	for n > 4 {
		n -= 3
		res *= 3
	}

	if n == 0 {
		return res
	}
	return res * n
}
