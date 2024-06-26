package loader

// https://leetcode.cn/problems/add-two-numbers/

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func tenPower(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

// TODO 这个是考察链表操作，首先得有HeadNode，然后后面的指针的调整

// case1:  直接求出合之后，再构建链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// first val
	var first, second int
	for i := 0; l1 != nil; i++ {
		first += l1.Val * (tenPower(i))
		l1 = l1.Next
	}

	for i := 0; l2 != nil; i++ {
		second += l2.Val * (tenPower(i))
		l2 = l2.Next
	}
	// add val
	sum := first + second
	var headList ListNode
	var lastIndex *ListNode
	for i := 0; ; i++ {
		lastval := sum % 10
		if lastIndex == nil {
			lastIndex = &ListNode{
				Val:  lastval,
				Next: nil,
			}
			headList.Next = lastIndex
		} else {
			listItem := &ListNode{
				Val:  lastval,
				Next: nil,
			}
			lastIndex.Next = listItem
			lastIndex = listItem
		}
		if sum/10 != 0 {
			sum = sum / 10
		} else {
			break
		}
	}
	return headList.Next
}

// case1:  双链表一直往前走

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	var headNode *ListNode
	var tailNode *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		if headNode == nil {
			newItem := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			headNode = newItem
			tailNode = newItem
		} else {
			newItem := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			tailNode.Next = newItem
			tailNode = newItem
		}
		carry = sum / 10
	}

	if carry > 0 {
		newItem := &ListNode{
			Val:  carry,
			Next: nil,
		}
		tailNode.Next = newItem
	}
	return headNode
}

func Test_addTwoNumbers(t *testing.T) {
	t.Logf("output is %s", reverseWords2("the sky is blue"))
}
