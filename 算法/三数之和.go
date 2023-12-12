package 算法

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return nil
	}

	sort.Ints(nums)

	n := len(nums)
	res := make([][]int, 0, n)
	for i, num := range nums[:n-2] {
		if i > 0 && num == nums[i-1] {
			continue
		}
		if num+nums[i+1]+nums[i+2] > 0 {
			break
		}

		j := i + 1
		k := n - 1
		for j < k {
			s := num + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				res = append(res, []int{num, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}

	return res
}
