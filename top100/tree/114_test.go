// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/?envType=study-plan-v2&envId=top-100-liked
package tree

func flatten(root *TreeNode) {
	// 	二叉树如何进行先序列遍历
	if root == nil {
		return
	}
	treeSlice := preOrderVisit(root)
	for i := 0; i != len(treeSlice)-1; i++ {
		treeSlice[i].Right = treeSlice[i+1]
		treeSlice[i].Left = nil
	}
	return
}

func preOrderVisit(root *TreeNode) []*TreeNode {
	var treeNodeSlice []*TreeNode
	if root == nil {
		return nil
	}
	treeNodeSlice = append(treeNodeSlice, root)
	// if root.Left != nil {
	leftSlice := preOrderVisit(root.Left)
	// }
	// if root.Right != nil {
	rightSlice := preOrderVisit(root.Right)
	// }

	treeNodeSlice = append(treeNodeSlice, leftSlice...)
	treeNodeSlice = append(treeNodeSlice, rightSlice...)
	return treeNodeSlice
}

// TODO: 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
//  1. 如果该节点的左子节点为空，则该节点不行进行展开操作
//  2. 如果该节点的左子节点不为空，则该节点的左子树的最后一个被访问的节点之后，
//     该节点的右子树被访问，（也就是右子树的前序节点）
//     所以问题就转化为找前序节点，然后然后调整，然后一步步展开，
//     总体还是有一点巧的方法
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root

	for cur != nil {
		next := cur.Left
		prev := next
		// find the most right
		for next != nil {
			prev = next
			next = next.Right
		}

		if prev != nil {
			prev.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
