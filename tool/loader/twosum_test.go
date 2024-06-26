package loader

// https://leetcode.cn/problems/two-sum/

// 方法一：暴力
func twoSum(nums []int, target int) []int {
	var index = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		index[0] = i
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				index[1] = j
				return index
			}
		}
	}
	return index
}

// 方法二：hash表

func twoSum2(nums []int, target int) []int {
	// TODO:  这里考察的是Map（hashmap）的操作和使用
	var hashmap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ifOk := hashmap[target-nums[i]]; ifOk {
			return []int{val, i}
		}
		hashmap[nums[i]] = i
	}
	return nil
}
