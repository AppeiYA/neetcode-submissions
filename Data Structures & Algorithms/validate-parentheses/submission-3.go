func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
    stack := []byte{}

	for i:=0; i < len(s); i++ {
		next := s[i]
		if len(stack) == 0 {
			stack = append(stack, next)
			continue
		}
		top := stack[len(stack)-1]
		if (top == '(' && next == ')') || 
			(top == '{' && next == '}') ||
			(top == '[' && next == ']') {
			stack = stack[:len(stack)-1]
		}else {
			stack = append(stack, next)
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
