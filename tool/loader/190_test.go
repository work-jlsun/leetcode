// https://leetcode.cn/problems/reverse-bits/
package loader

// 今天的题目是要求将一个数字，把其二进制翻转，求得到的另外一个二进制数。
// 基本的思路，循环移位
func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i <= 31; i++ {
		ret = (ret << 1) | (num & 1)
		num = num >> 1
	}
	return ret
}
