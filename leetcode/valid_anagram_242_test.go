package leetcode

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {

	s1 := "anagram"
	s2 := "nagaram"

	flag := isAnagram(s1, s2)
	fmt.Printf("isAnagram(%v, %v) = %v\n", s1, s2, flag)
}
