package tree

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/?envType=study-plan-v2&envId=top-100-liked

// 这个问题关键点：在于递归的当返回结果超过大小之后，就返回
func kthSmallest(root *TreeNode, k int) int {
	// 中旬遍历
	intSlice := inorderTraverse2(root, k)

	if len(intSlice) >= k {
		return intSlice[k-1]
	}
	return 0
}

func inorderTraverse2(root *TreeNode, k int) []int {
	var intSlice []int
	if root == nil {
		return nil
	}
	if root.Left != nil {
		intSlice = append(intSlice, inorderTraverse2(root.Left, k)...)
	}
	intSlice = append(intSlice, root.Val)
	if len(intSlice) >= k {
		return intSlice
	}
	if root.Right != nil {
		intSlice = append(inorderTraverse2(root.Right, k), intSlice...)
	}
	return intSlice
}
