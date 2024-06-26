package loader

// https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/?envType=study-plan-v2&envId=leetcode-75
func kidsWithCandies(candies []int, extraCandies int) []bool {

	max := func() int {
		max := candies[0]
		for i := 0; i < len(candies); i++ {
			if candies[i] > max {
				max = candies[i]
			}
		}
		return max
	}

	maxnum := max()
	ret := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxnum {
			ret[i] = true
		} else {
			ret[i] = false
		}
	}
	return ret
}
