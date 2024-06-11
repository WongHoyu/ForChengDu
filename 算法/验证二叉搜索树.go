package 算法

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return Dfs(root, math.MinInt, math.MaxInt)
}

func Dfs(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return Dfs(node.Left, min, node.Val) && Dfs(node.Right, node.Val, max)
}
