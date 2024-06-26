// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
// 105. 从前序与中序遍历序列构造二叉树
package tree

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	leftNode := root
	rightNode := root
	for i := 1; i < len(preorder); i++ {
		num := preorder[i]
		index := findIndex(num, inorder)
		if index < 0 {
			return nil
		}
		if index < i {
			leftNode.Left = &TreeNode{
				Val:   preorder[i],
				Left:  nil,
				Right: nil,
			}
			leftNode = leftNode.Left
		} else {
			rightNode.Right = &TreeNode{
				Val:   preorder[i],
				Left:  nil,
				Right: nil,
			}
			rightNode = rightNode.Right
		}
	}
	return nil
}

func findIndex(n int, inorder []int) int {
	for i, num := range inorder {
		if num == n {
			return i
		}
	}
	return -1
}
