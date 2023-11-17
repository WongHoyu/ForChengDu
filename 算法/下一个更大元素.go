package 算法

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	stack := make([]int, 0, len(nums2))

	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		for len(stack) > 0 && n > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			mp[n] = stack[len(stack)-1]
		} else {
			mp[n] = -1
		}

		stack = append(stack, n)
	}

	ans := make([]int, 0, len(nums1))
	for _, n := range nums1 {
		ans = append(ans, mp[n])
	}

	return ans
}
