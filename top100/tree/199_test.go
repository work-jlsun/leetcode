package tree

// https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked

// 右看图，藐视就是层级遍历的最有一个节点
func rightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	var level int = 0
	var intResult []int

	var treeNodeQueue []*TreeNode
	treeNodeQueue = append(treeNodeQueue, root)
	for len(treeNodeQueue) != 0 {
		level++
		var tmp []*TreeNode
		// add the right result
		levelLength := len(treeNodeQueue)
		intResult = append(intResult, treeNodeQueue[levelLength-1].Val)
		for i := 0; i != levelLength; i++ {
			if treeNodeQueue[i].Left != nil {
				tmp = append(tmp, treeNodeQueue[i].Left)
			}
			if treeNodeQueue[i].Right != nil {
				tmp = append(tmp, treeNodeQueue[i].Right)
			}
		}
		treeNodeQueue = tmp
	}

	return intResult
}
