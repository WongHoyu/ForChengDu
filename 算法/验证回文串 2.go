package 算法

func validPalindrome(s string) bool {
	if s == "" {
		return true
	}

	f := 0
	b := len(s) - 1
	for f < b {
		if s[f] != s[b] {
			return valid(s[f:b]) || valid(s[f+1:b+1])
		}

		f++
		b--
	}

	return true
}

func valid(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return valid(s[1 : len(s)-1])
}
