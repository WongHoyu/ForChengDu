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

	return BST(root, math.MinInt, math.MaxInt)
}

func BST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if min >= node.Val || max <= node.Val {
		return false
	}

	return BST(node.Left, min, node.Val) && BST(node.Right, node.Val, max)
}
