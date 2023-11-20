package 算法

func firstUniqChar(s string) int {
	exist := make(map[int32]int, 26)
	for i, v := range s {
		_, ok := exist[v]
		if ok {
			exist[v] = -1
		} else {
			exist[v] = i
		}
	}

	for _, v := range s {
		if i := exist[v]; i != -1 {
			return i
		}
	}

	return -1
}
