package 算法

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0

	result := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if rightMax > leftMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}

	return result
}
