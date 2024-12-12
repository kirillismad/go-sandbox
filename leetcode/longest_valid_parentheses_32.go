package leetcode

// https://leetcode.com/problems/longest-valid-parentheses/

// #stack

func longestValidParentheses(s string) int {
	var result int
	stack := []int{-1}
	for i, r := range []rune(s) {
		if r == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			result = max(result, i-stack[len(stack)-1])
		}
	}
	return result
}
