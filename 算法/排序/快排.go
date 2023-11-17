package 排序

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	p := partition(nums, left, right)

	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func partition(nums []int, left int, right int) int {
	p := nums[left]

	for left < right {
		for left < right && nums[right] >= p {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= p {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = p
	return left
}
