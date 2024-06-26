package tree

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
func maxDepth(root *TreeNode) int {
	leftDepth := 0
	rightDepth := 0
	if root == nil {
		return 0
	}
	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
