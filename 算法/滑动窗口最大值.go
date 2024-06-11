package 算法

func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0, k+1)
	result := make([]int, 0, len(nums))

	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if i-q[0] >= k {
			q = q[1:]
		}

		if i >= k-1 {
			result = append(result, nums[q[0]])
		}
	}

	return result
}
