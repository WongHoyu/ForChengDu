package 算法

func firstUniqChar(s string) int {
	exist := make(map[int32]int, 26)
	for _, c := range s {
		exist[c]++
	}

	for i, c := range s {
		if exist[c] == 1 {
			return i
		}
	}

	return -1
}
