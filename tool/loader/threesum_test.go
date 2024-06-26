package loader

import (
	"sort"
	"testing"
)

// https://leetcode.cn/problems/3sum/

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

//  [-4, -1, -1, 0, 1, 2]

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

//  [-4, -1, -1, 0, 1, 2]
// 如何进行去重呢?

func threeSum(nums []int) [][]int {
	var result [][]int
	// 注意slice不可比较，array是可以比较的
	var m = make(map[[3]int]bool)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// 这个减枝注意一下，相同数字情况下，肯定已经包含后续的情况了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1

		// 这里注意下，由于是排序的，采取包围策略即可
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res := [3]int{nums[i], nums[j], nums[k]}

				// 还有就是这里的map比对三元组去重的策略
				if _, exist := m[res]; exist {
					continue
				}
				result = append(result, res[:])
				m[res] = true
				j++
				k--
			}
			if sum > 0 {
				k--
			}
			if sum < 0 {
				j++
			}
		}
	}
	return result
}

func Test_ThreeNumbers(t *testing.T) {
	t.Logf("output is %+v", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// 如何进行去重呢?
//
// func threeSum(nums []int) [][]int {
// 	var result [][]int = make([][]int, 0)
// 	sort.Ints(nums)
// 	for i := 0; i < len(nums)-3; i++ {
// 		firstElem := nums[i]
//
//
// 		for secondIndex := i + 1; secondIndex < len(nums)-2; secondIndex++ {
// 			secondElem := nums[secondIndex]
// 			for lastIndex := len(nums) - 1; lastIndex > secondElem+1; lastIndex-- {
// 				lastElem := nums[lastIndex]
// 				sum := firstElem + secondElem + lastElem
// 				if sum == 0 {
// 					result = append(result, []int{firstElem, secondElem, lastElem})
// 				}
// 				if sum < 0 {
// 					break
// 				}
// 				if sum > 0 {
// 					continue
// 				}
// 			}
// 		}
// 	}
// 	return result
// }
