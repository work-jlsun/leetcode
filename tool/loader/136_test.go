// https://leetcode.cn/problems/single-number/

package loader

// 这里考察的事 位操作
func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}
