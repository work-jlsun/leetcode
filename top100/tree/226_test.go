package tree

// https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-100-liked
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	var left *TreeNode
	var right *TreeNode
	if root == nil {
		return root
	}

	if root.Left != nil {
		left = invertTree(root.Left)
	}

	if root.Right != nil {
		right = invertTree(root.Right)
	}

	root.Left = right
	root.Right = left
	// 注意是范围root
	return root
}
