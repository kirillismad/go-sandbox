package leetcode

import "unicode"

// https://leetcode.com/problems/valid-palindrome/
// #two-pointers

func isPalindrome(s string) bool {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1
	for left < right {
		switch {
		case !(unicode.In(runes[left], unicode.Letter, unicode.Digit)):
			left++
		case !(unicode.In(runes[right], unicode.Letter, unicode.Digit)):
			right--
		case unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]):
			return false
		default:
			left++
			right--
		}
	}
	return true
}
