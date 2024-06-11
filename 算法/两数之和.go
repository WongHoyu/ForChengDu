package 算法

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))
	for i, num := range nums {
		t := target - num
		if _, ok := tmp[t]; ok {
			return []int{tmp[t], i}
		}
		tmp[num] = i
	}

	return []int{}
}
