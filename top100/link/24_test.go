package link

// 方案：3指针法
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	newHead := &ListNode{
		Next: head,
	}

	pre := newHead
	first := head
	second := first.Next
	for !(first == nil || second == nil) {

		//  swap
		secondNext := second.Next
		pre.Next = second
		second.Next = first
		first.Next = secondNext

		//
		pre = first
		first = first.Next
		if first == nil {
			break
		}
		second = first.Next
	}
	return newHead.Next
}
