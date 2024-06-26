package loader

import "testing"

// https://leetcode.cn/problems/longest-palindromic-substring/solutions/1348874/s-by-xext-5zk3

// 5. 最长回文子串 ： 中心扩散法
func longestPalindrome(s string) string {
	var startIndex int
	var maxLength int
	for i := 0; i < len(s)-1; i++ {
		oddMaxLength := expand(s, i, i)
		evenMaxLength := expand(s, i, i+1)

		maxLen := func() int {
			if oddMaxLength > evenMaxLength {
				return oddMaxLength
			} else {
				return evenMaxLength
			}
		}()

		if maxLen > maxLength {
			maxLength = maxLen
			if oddMaxLength >= evenMaxLength {
				startIndex = i - oddMaxLength/2
			} else {
				startIndex = i - oddMaxLength/2
			}
		}
	}
	return s[startIndex : startIndex+maxLength]
}

func expand(s string, start int, end int) int {
	tmpLen := end - start + 1
	if s[start] != s[end] {
		return 0
	}
	var equals int
	for {
		start = start - 1
		end = end + 1
		if start < 0 || end >= len(s) {
			break
		}
		if s[start] == s[end] {
			equals++
		}
	}
	maxLength := 2*equals + tmpLen
	return maxLength
}

// 方法二：动态规划（暂时不深入）

func Test_longestPalindrome(t *testing.T) {
	t.Logf("output is %s", longestPalindrome("babad"))
}
