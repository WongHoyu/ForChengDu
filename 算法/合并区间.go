package 算法

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[1] < intervals[i][0] {
			res = append(res, prev)
			prev = intervals[i]
		} else {
			prev[1] = max(prev[1], intervals[i][1])
		}
	}

	res = append(res, prev)
	return res
}
