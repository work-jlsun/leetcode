package stack

func isValid(s string) bool {
	var element = make([]byte, len(s))
	var top int = -1
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if top == -1 {
			if s[i] == '{' || s[i] == '[' || s[i] == '(' {
				top++
				element[top] = s[i]
				continue
			} else {
				return false
			}
		} else {
			if element[top] == '{' && s[i] == '}' {
				top--
				continue
			} else if element[top] == '(' && s[i] == ')' {
				top--
				continue
			} else if element[top] == '[' && s[i] == ']' {
				top--
				continue
			} else {
				return false
			}
		}
	}
	return true
}
