package tree

// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ints []int
	if root == nil {
		return ints
	}
	if root.Left != nil {
		ints = append(inorderTraversal(root.Left), ints...)
	}
	ints = append(ints, root.Val)
	if root.Right != nil {
		ints = append(ints, inorderTraversal(root.Right)...)
	}
	return ints
}
