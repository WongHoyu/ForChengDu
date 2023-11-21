package 算法

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	exist := make(map[int]bool, len(nums1))
	for _, v := range nums1 {
		exist[v] = true
	}

	ans := make([]int, 0, len(nums1))
	for _, v := range nums2 {
		if exist[v] {
			ans = append(ans, v)
			exist[v] = false
		}
	}

	return ans
}
