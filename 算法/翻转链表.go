package 算法

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead *ListNode
	for head != nil {
		p := head.Next
		head.Next = newHead
		newHead = head
		head = p
	}

	return newHead
}
