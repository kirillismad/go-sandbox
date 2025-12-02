package leetcode

// https://leetcode.com/problems/valid-parentheses
func isValid(s string) bool {
	var closeOpen = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	push := func(r rune) {
		stack = append(stack, r)
	}
	peek := func() rune {
		return stack[len(stack)-1]
	}

	pop := func() rune {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return r
	}
	empty := func() bool {
		return len(stack) == 0
	}

	for _, r := range s {
		open, isClosing := closeOpen[r]
		if isOpening := !isClosing; isOpening {
			push(r)
			continue
		}
		if empty() || peek() != open {
			return false
		}
		pop()
	}
	return empty()
}
