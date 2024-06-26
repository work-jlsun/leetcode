package heap

import "sort"

// 方案1：各heap，heap， 建1个heap不会，放弃
func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
	return nums[len(nums)-k]
}

//
