package link

func hasCycle(head *ListNode) bool {

	var nodeMap = make(map[*ListNode]bool)

	if head == nil || head.Next == nil {
		return true
	}
	for head != nil {
		if nodeMap[head] {
			return true
		}
		nodeMap[head] = true
	}
	return false
}

// 快慢指针问题，进入换之后，快指针一定会追上慢指针
// 只是追上的形式有2种，一个是差2不正好重合，还有1个是差一步跳过
// 然后下一步慢指针走一步重合
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for {
		if slow == fast {
			return true
		}

		// step slow
		if slow.Next == nil {
			return false
		}
		slow = slow.Next
		if slow == fast {
			return true
		}

		// step fast
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
}
