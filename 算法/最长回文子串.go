package 算法

func longestPalindrome(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		l := i
		r := i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(result) {
			result = s[l+1 : r]
		}

		l = i
		r = i + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(result) {
			result = s[l+1 : r]
		}
	}

	return result
}
