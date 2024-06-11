package 算法

// s = "abc"
func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	result := 0
	for i := 0; i < len(s); i++ {
		l := i
		r := i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		l = i
		r = i + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}
