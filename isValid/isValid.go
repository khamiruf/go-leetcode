func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pair := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := []rune{}

	for _, r := range s {
		if _, ok := pair[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pair[stack[len(stack)-1]] != r {
			return false
		} else {
			// pop the stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}