package 算法

func mergeTwoLists(node1 *ListNode, node2 *ListNode) *ListNode {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	}

	if node1.Val <= node2.Val {
		node1.Next = mergeTwoLists(node1.Next, node2)
		return node1
	} else {
		node2.Next = mergeTwoLists(node1, node2.Next)
		return node2
	}
}
