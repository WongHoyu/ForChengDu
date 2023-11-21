package 算法

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for {
			if strings.HasPrefix(strs[i], prefix) {
				break
			}
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}
