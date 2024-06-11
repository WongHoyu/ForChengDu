package 算法

func buildTree(preorder []int, inorder []int) *TreeNode {
	for i, v := range inorder {
		if preorder[0] == v {
			return &TreeNode{
				Val:   v,
				Left:  buildTree(preorder[1:i+1], inorder[0:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}
