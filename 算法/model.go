package 算法

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
