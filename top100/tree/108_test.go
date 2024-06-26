package tree

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked
func sortedArrayToBST(nums []int) *TreeNode {

	if nums == nil && len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		node := &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
		return node
	}
	if len(nums) == 2 {
		node2 := &TreeNode{
			Val:   nums[1],
			Left:  nil,
			Right: nil,
		}
		node1 := &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
		node2.Left = node1
		return node2
	}

	middleIndex := len(nums) / 2

	middleNode := &TreeNode{
		Val:   nums[middleIndex],
		Left:  nil,
		Right: nil,
	}
	middleNode.Left = sortedArrayToBST(nums[:middleIndex])
	middleNode.Right = sortedArrayToBST(nums[middleIndex+1:])
	return middleNode
}
