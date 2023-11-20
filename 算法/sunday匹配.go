package 算法

func strStr(str1 string, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}
	if len(str2) > len(str1) {
		return -1
	}
	if str2 == str1 {
		return 0
	}

	result := -1
	i := 0
	matchIndex := 0

	for i < len(str1) {
		if str1[i] == str2[matchIndex] {
			for str1[i] == str2[matchIndex] {
				if result == -1 {
					result = i
				}

				i++
				matchIndex++

				if matchIndex >= len(str2) {
					return result
				}
			}
		} else {
			result = -1
			i++
		}
	}

	return result
}
