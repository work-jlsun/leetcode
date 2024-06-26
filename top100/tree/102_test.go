// https://leetcode.cn/problems/binary-tree-level-order-traversal/?envType=study-plan-v2&envId=top-100-liked
package tree

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	var result [][]int

	var levelQueue []*TreeNode
	levelQueue = append(levelQueue, root)

	level := 0

	// 层序遍历的关键，是第一层，然后创建一个第2层的slice，知道所有的层处理完
	for len(levelQueue) > 0 {
		var tempLevelQueue []*TreeNode
		result = append(result, []int{})
		for i := 0; i < len(levelQueue); i++ {
			node := levelQueue[i]
			result[level] = append(result[level], node.Val)
			if node.Left != nil {
				tempLevelQueue = append(tempLevelQueue, node.Left)
			}
			if node.Right != nil {
				tempLevelQueue = append(tempLevelQueue, node.Right)
			}
		}
		levelQueue = tempLevelQueue
		level++
	}
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {
		ret = append(ret, []int{})
		// temp for next level，关键在这个下一层
		tmp := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			ret[i] = append(ret[i], node.Val)
			if queue[j].Right != nil {
				tmp = append(tmp, node.Right)
			}
			if queue[j].Left != nil {
				tmp = append(tmp, node.Left)
			}
		}
		// 遍历下一层
		queue = tmp
	}
	return ret
}
