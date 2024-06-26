package link

// https://leetcode.cn/problems/sort-list/description/?envType=study-plan-v2&envId=top-100-liked
import "sort"

// 这里说白了就是直接xx
func sortList(head *ListNode) *ListNode {

	var nodeSlice []*ListNode

	for head != nil {
		nodeSlice = append(nodeSlice, head)
		head = head.Next
	}

	// 直接
	sort.Slice(nodeSlice, func(i, j int) bool {
		if nodeSlice[i].Val < nodeSlice[j].Val {
			return true
		}
		return false
	})

	for i := 0; i < len(nodeSlice); i++ {
		if i == (len(nodeSlice) - 1) {
			nodeSlice[i].Next = nil
		} else {
			nodeSlice[i].Next = nodeSlice[i+1]
		}
	}

	return nodeSlice[0]
}
