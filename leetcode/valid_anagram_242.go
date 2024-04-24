package leetcode

// https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	runeCounter := make(map[rune]int)

	for _, r := range s {
		runeCounter[r]++
	}

	for _, r := range t {
		runeCounter[r]--
	}

	for _, counter := range runeCounter {
		if counter != 0 {
			return false
		}
	}

	return true
}
