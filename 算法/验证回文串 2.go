package 算法

func validPalindrome(s string) bool {
	if s == "" {
		return true
	}

	count := 1
	f := 0
	b := len(s) - 1
	for f < b {
		if s[f] != s[b] {
			if count == 1 {
				if s[f+1] == s[b] {
					f++
				} else if s[f] == s[b-1] {
					b--
				} else {
					return false
				}

				count--
				continue
			}

			return false
		}

		f++
		b--
	}

	return true
}
