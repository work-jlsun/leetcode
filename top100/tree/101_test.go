// https://leetcode.cn/problems/symmetric-tree/solutions/46560/dong-hua-yan-shi-101-dui-cheng-er-cha-shu-by-user7/?envType=study-plan-v2&envId=top-100-liked
package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归怎么判断轴对称
// 退出条件为递归到底部?
// 左边递归到3
// 右边 对轨道3
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return compare(root.Right, root.Left)
}

func compare(right *TreeNode, left *TreeNode) bool {

	if right == nil && left == nil {
		return true
	}
	if right == nil || left == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
