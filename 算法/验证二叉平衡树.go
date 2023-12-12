package 算法

func isBalanced(root *TreeNode) bool {
	return getHigh(root) != -1
}

func getHigh(node *TreeNode) int {
	if node == nil {
		return 0
	}

	lHigh := getHigh(node.Left)
	if lHigh == -1 {
		return -1
	}

	rHigh := getHigh(node.Right)
	if rHigh == -1 {
		return -1
	}

	if getAbs(lHigh-rHigh) > 1 {
		return -1
	}

	return getMax(lHigh, rHigh) + 1
}
