package link

// 方案一：  O(n) 时间复杂度和 O(n) 空间复杂度
func isPalindrome(head *ListNode) bool {
	var intSlice []int
	for head != nil {
		intSlice = append(intSlice, head.Val)
		head = head.Next
	}
	for i := 0; i < len(intSlice)/2; i++ {
		if intSlice[i] != intSlice[len(intSlice)-i-1] {
			return false
		}
	}
	return true
}

// 方案二：  O(n) 时间复杂度和 O(1) 空间复杂度， 递归的解法

// 首先，递归到尾部，然后进行比较
// 然后返回进行比较
func isPalindrome2(head *ListNode) bool {

	startCmpIndex := head

	var comFunc func(node *ListNode) bool
	comFunc = func(node *ListNode) bool {
		if node != nil {
			// 如果不是空节点，比较值是否相等，如果不等范围 false,这样递归一直返回false退出
			if !comFunc(node.Next) {
				return false
			}
			if node.Val != startCmpIndex.Val {
				return false
			}
			startCmpIndex = startCmpIndex.Next
			return true
		}
		// 遇到空节点返回
		return true
	}
	return comFunc(head)
}
