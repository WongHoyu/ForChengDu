package 算法

func lengthOfLongestSubstring(s string) int {
	cache := make(map[int32]int)
	left := 0
	ans := 0

	for right, c := range s {
		cache[c] += 1
		for cache[c] > 1 {
			cache[int32(s[left])] -= 1
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
