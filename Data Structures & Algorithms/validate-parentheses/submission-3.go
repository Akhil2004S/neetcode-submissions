func isValid(s string) bool {
    	var stack []rune
	if len(s)%2 != 0 {
		return false
	}
	for _, char := range s {
		if char == '}' || char == ']' || char == ')' {
			if len(stack) == 0 {
				// Underflow
				return false
			}
			top := stack[len(stack)-1]
			if (top == '{' && char == '}') || (top == '(' && char == ')') || (top == '[' && char == ']') {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
		if len(stack) == len(s) {
			// Overflow
			return false
		}
		stack = append(stack, char)
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
