package loader

import (
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/reverse-words-in-a-string/description/

// case1 use strings split
func reverseWords(s string) string {
	var ret string
	strs := strings.Split(s, " ")
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != " " {
			ret = ret + " " + strs[i]
		}
	}
	return ret
}

func reverseWords2(s string) string {
	var words []string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else {
			j := i
			for ; j < len(s) && s[j] != ' '; j++ {
				// noop
				fmt.Printf("%d\n", j)
			}
			words = append(words, s[i:j])
			i = j
		}
	}
	var ret string
	for i := len(words) - 1; i >= 0; i-- {
		if ret != "" {
			ret = ret + " " + words[i]
		} else {
			ret = words[i]
		}
	}
	return ret
}

func Test_reverseWords(t *testing.T) {
	t.Logf("output is %s", reverseWords2("the sky is blue"))
}
