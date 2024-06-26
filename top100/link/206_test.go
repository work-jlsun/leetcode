package link

// https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方案1
func reverseList2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	cur := head
	curNext := head.Next

	for cur.Next != nil {
		cur.Next = prev
		prev = cur
		cur = curNext
		curNext = cur.Next
	}
	cur.Next = prev
	return cur
}

// 方案2: 更加直观的方案
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		// 防止断掉，保存下一个节点，然后调整指针
		next := cur.Next
		cur.Next = prev

		// 调整位置，志祥下一个待处理的节点，并设置新的prev
		prev = cur
		cur = next
	}
	return prev
}
