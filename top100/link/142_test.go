package link

// https://leetcode.cn/problems/linked-list-cycle-ii/description/?envType=study-plan-v2&envId=top-100-liked
func detectCycle(head *ListNode) *ListNode {

	var m = make(map[*ListNode]int)
	next := head
	for i := 0; next != nil; i++ {
		if _, existed := m[next]; existed {
			return next
		}
		m[next] = next.Val
		next = next.Next
	}
	return nil
}
