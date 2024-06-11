package 算法

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	str = strings.TrimSpace(str)

	sign := 1
	res := 0
	for i, s := range str {
		if s > '0' && s <= '9' {
			res = res*10 + int(s-'0')
		} else if s == '-' && i == 0 {
			sign = -1
		} else if s == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if res > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * res
}
