package 算法

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, len(nums)-k)
	return nums[len(nums)-k]
}

func quickSort(nums []int, left int, right int, k int) {
	if left >= right {
		return
	}

	p := partition(nums, left, right)

	if k == p {
		return
	} else if k < p {
		quickSort(nums, left, p-1, k)
	} else {
		quickSort(nums, p+1, right, k)
	}
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
