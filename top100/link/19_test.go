package link

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-100-liked

// 方案1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// calc the length
	moveHead := head
	var length int
	for moveHead != nil {
		length++
		moveHead = moveHead.Next
	}

	// case1: delete the head
	if length == n {
		return head.Next
	}

	// case2: delete the middle
	var dumpHead = &ListNode{
		Next: head,
	}
	for i := 0; i < length-n; i++ {
		dumpHead = dumpHead.Next
	}
	dumpHead.Next = dumpHead.Next.Next
	return head
}

// 方案二： 堆栈
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var stack = make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	if n == len(stack) {
		if len(stack) > 1 {
			return stack[1]
		} else {
			return nil
		}
	} else {
		stack[len(stack)-n-1].Next = stack[len(stack)-n-1].Next.Next
		return stack[0]
	}
}

// 方案三: TODO 双指针法
