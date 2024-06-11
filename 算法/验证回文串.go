package 算法

import "strings"

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}

	s = strings.ToLower(s)

	l := 0
	r := len(s) - 1
	for l < r {
		for l < r && s[l]-'0' > 9 && s[l]-'a' > 25 {
			l++
		}
		for l < r && s[r]-'0' > 9 && s[r]-'a' > 25 {
			r--
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}
	return true
}
