package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.cn/problems/validate-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked

// 二叉搜索树，本质上是终须遍历结果是正确的
func isValidBST(root *TreeNode) bool {
	inorderSlice := inorderTraversal(root)
	if inorderSlice == nil {
		return true
	}
	for i := 0; i < len(inorderSlice)-1; i++ {
		if inorderSlice[i] <= inorderSlice[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func inorderTraversal2(root *TreeNode) []int {
	var intSlice []int
	if root == nil {
		return nil
	}
	if root.Left != nil {
		intSlice = append(inorderTraversal2(root.Left), intSlice...)
	}
	intSlice = append(intSlice, root.Val)
	if root.Right != nil {
		intSlice = append(intSlice, inorderTraversal2(root.Right)...)
	}
	return intSlice
}
