package 算法

func lengthOfLongestSubstring(s string) int {
	res := 0

	cache := make(map[int32]int, 128)
	left := 0
	for right, c := range s {
		cache[c] += 1
		for cache[c] > 1 {
			cache[int32(s[left])] -= 1
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}
