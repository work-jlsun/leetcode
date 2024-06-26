package link

// https://leetcode.cn/problems/linked-list-cycle-ii/description/?envType=study-plan-v2&envId=top-100-liked
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head = &ListNode{}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		if list1.Val <= list2.Val {
			head.Next = list1
		} else {
			head.Next = list2
		}
	}
	runHead := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			runHead.Next = list2
			break
		}
		if list2 == nil {
			runHead.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			runHead.Next = list1
			list1 = list1.Next
			continue
		}
		if list1.Val > list2.Val {
			runHead.Next = list2
			list2 = list2.Next
			continue
		}
	}
	return head.Next
}
