package leetcode

// https://leetcode.com/problems/valid-parentheses
func isValid(s string) bool {
	var closeOpen = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune

	for _, r := range s {
		open, isClosing := closeOpen[r]
		if !isClosing {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
