package 算法

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	var res int
	var minus bool
	for _, v := range str {
		if (v-'0' >= 10 || v-'0' < 0) && v != '-' {
			if res != 0 {
				break
			}
			continue
		}
		if v == '-' {
			if minus {
				break
			}
			minus = true
			continue
		}

		res = res*10 + int(v-'0')
	}

	if !minus {
		if res > (2<<30 - 1) {
			return 2<<30 - 1
		}
		return res
	}

	if res > 2<<30 {
		return -1 * (2 << 30)
	}

	return res * -1
}
