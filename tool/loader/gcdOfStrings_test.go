package loader

import "testing"

func gcdOfStrings(str1 string, str2 string) string {

	len1, len2 := len(str1), len(str2)
	max := func() int {
		if len1 < len2 {
			return len1
		}
		return len2
	}

	for i := max(); i > 0; i-- {
		if len1%i == 0 && len2%i == 0 {
			if check(str1+str2, str1[0:i]) {
				return str1[0:i]
			}
		}
	}
	return ""
}

func check(s string, sub string) bool {
	for i := 0; i < len(s)/len(sub); i++ {
		str := s[i*len(sub) : (i+1)*len(sub)]
		if str != sub {
			return false
		}
	}
	return true
}

func Test_a(t *testing.T) {
	t.Logf("output is %s", gcdOfStrings("ABCABC", "ABC"))
}
