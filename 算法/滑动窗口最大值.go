package 算法

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	q := make([]int, 0, k+1)
	result := make([]int, 0, len(nums))
	for i, num := range nums {
		// 保持队列单调性
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}
		// 入队
		q = append(q, i)

		// 队头已经超出滑动窗口范围
		if i-q[0] >= k {
			q = q[1:]
		}

		// 下标 k-1 开始记录
		if i >= k-1 {
			result = append(result, nums[q[0]])
		}
	}

	return result
}
