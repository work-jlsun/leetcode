// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
package link

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方案1： 使用直观的放哪，直接 hash索引的方式查找
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	// 1. slice array a1
	// 2. slice arry a2
	// 3. compare from the end
	// 4. map insert a
	// 5. try to insert b,if noexist

	var m = make(map[*ListNode]bool)

	head := headA
	for head != nil {
		m[head] = true
		head = head.Next
	}

	head = headB
	for head != nil {
		if m[head] {
			return head
		}
		head = head.Next
	}
	return nil
}

// 方案2：假设有交点，那么将A的尾部与B项链，B的尾部与A相连的遍历方式，2者总会相遇
