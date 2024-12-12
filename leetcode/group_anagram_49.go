package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/group-anagrams/

// #hashmap

func groupAnagrams(strs []string) [][]string {
	keyAnagrams := make(map[string][]string)
	result := make([][]string, 0)

	for i := 0; i < len(strs); i++ {
		key := []rune(strs[i])
		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})
		keyAnagrams[string(key)] = append(keyAnagrams[string(key)], strs[i])
	}

	for _, v := range keyAnagrams {
		result = append(result, v)
	}
	return result
}
