package 算法

import "strings"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)

	// aba
	// aa
	f := 0
	b := len(s) - 1
	for f < b {
		for f < b && s[f]-'0' > 9 && s[f]-'a' > 25 {
			f++
		}
		for f < b && s[b]-'0' > 9 && s[b]-'a' > 25 {
			b--
		}

		if s[f] != s[b] {
			return false
		}

		f++
		b--
	}

	return true
}
