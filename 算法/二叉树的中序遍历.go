package 算法

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	BST2(root, &result)

	return result
}

func BST2(node *TreeNode, treeList *[]int) {
	if node == nil {
		return
	}

	BST2(node.Left, treeList)
	*treeList = append(*treeList, node.Val)
	BST2(node.Right, treeList)
}
