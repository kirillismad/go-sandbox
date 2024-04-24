package examples

import (
	"fmt"
	"sort"
)

func DemoSlices() {
	// sliceFromArray()
	// changesInSlices()
	// initSlice()
	// appendInSlice()
	// extendSlice()
	// deleteInSlice()
	// copySlice()
	// twoDimensionSlices()
	// shrinkSlice()
	// sortReverseOrder()
	// sortUsingFewFields()
	demoDirtyFuncWithSlice()
}

func sliceFromArray() {
	fmt.Println(newline + fname(sliceFromArray))
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr[:] =", arr[:])

	for i := 1; i < len(arr); i++ {
		fmt.Printf("arr[%v:] = %v\n", i, arr[i:])
	}

	for i := len(arr); i > 0; i-- {
		fmt.Printf("arr[:%v] = %v\n", i, arr[:i])
	}

	for i := 0; i <= int(len(arr)/2); i++ {
		right := len(arr) - i
		fmt.Printf("arr[%v:%v] = %v\n", i, right, arr[i:right])
	}
}

func changesInSlices() {
	fmt.Println(newline + fname(changesInSlices))
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr =", arr)

	s1 := arr[1:4] // 2,3,4
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n", s1, len(s1), cap(s1))

	s2 := s1[1:2] // 3,4
	fmt.Printf("s2 = %v, len(s2) = %v, cap(s2) = %v\n", s2, len(s2), cap(s2))

	s1[1] = 99
	fmt.Println("arr =", arr)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

	arr[2] = 88
	fmt.Println("arr =", arr)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

	s2[0] = 77
	fmt.Println("arr =", arr)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

}

func initSlice() {
	fmt.Println(newline + fname(initSlice))

	var nilSlice []int

	if nilSlice == nil {
		fmt.Println("nilSlice =", nilSlice)
		// x := nilSlice[0] // panic
		nilSlice = append(nilSlice, 42)
		fmt.Println("nilSlice =", nilSlice)
	}

	s1 := make([]int, 3) // (type, len)
	fmt.Println("s1 =", s1)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 =", s2)

	s3 := make([]int, 3, 6) // (type, len, cap)
	fmt.Println("s3 =", s3)

	s4 := []int{}
	fmt.Println("s4 =", s4)
}

func appendInSlice() {
	fmt.Println(newline + fname(appendInSlice))
	s := make([]int, 0)
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n\n", s, len(s), cap(s))

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))
	}

}

func extendSlice() {
	fmt.Println(newline + fname(extendSlice))
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	s := append(s1, s2...)
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))
}

func deleteInSlice() {
	fmt.Println(newline + fname(deleteInSlice))
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))

	const index = 2
	s = append(s[:index], s[index+1:]...)
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))

	s = s[:cap(s)] // side effect if extending
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))
}

func copySlice() {
	fmt.Println(newline + fname(copySlice))
	s := []int{1, 2, 3}
	fmt.Printf("s = %v, len(s) = %v, cap(s) = %v\n", s, len(s), cap(s))

	c := make([]int, len(s))
	copy(c, s)
	fmt.Printf("c = %v, len(c) = %v, cap(c) = %v\n", c, len(c), cap(c))

	c[0] = 11
	c = append(c, 44)
	fmt.Printf("c = %v, len(c) = %v, cap(c) = %v\n", c, len(c), cap(c))
}

func twoDimensionSlices() {
	fmt.Println(newline + fname(twoDimensionSlices))
	const l = 3
	s := make([][]int, l)
	fmt.Println("s =", s)

	for i := 0; i < l; i++ {
		for j := 0; j < i+1; j++ {
			s[i] = append(s[i], j+i*l)
		}
	}
	fmt.Println("s =", s)
}

func shrinkSlice() {
	s := []int{1, 2, 3}
	s = s[:0]
	fmt.Printf("s: %v\n", s)
}

func sortReverseOrder() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

type s1 struct {
	id         int
	name       string
	percentage float64
}

type s1Slice []s1

func (s s1Slice) Len() int {
	return len(s)
}

func (s s1Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s s1Slice) Less(i, j int) bool {
	switch {
	case s[i].id != s[j].id:
		return s[i].id < s[j].id
	case s[i].name != s[j].name:
		return s[i].name < s[j].name
	default:
		return s[i].percentage < s[j].percentage
	}
}
func randomS1() s1 {
	return s1{
		id:         rint(100),
		name:       rstring(),
		percentage: rfloat(100),
	}
}

func sortUsingFewFields() {
	s := make(s1Slice, 10)
	for i := range s {
		s[i] = randomS1()
	}

	sort.Sort(s)
	for i, v := range s {
		fmt.Printf("%v) %#v\n", i, v)
	}
	fmt.Println("---------------------------")

	sort.Sort(sort.Reverse(s))
	for i, v := range s {
		fmt.Printf("%v) %#v\n", i, v)
	}
}

func dirtySliceFunc(s []int) {
	s = append(s, 456)
}

func demoDirtyFuncWithSlice() {
	s := make([]int, 0, 3)
	s = append(s, 1)
	s = append(s, 2)
	dirtySliceFunc(s)
	fmt.Printf("s: %v\n", s)
	s = s[:cap(s)]
	fmt.Println(s)
}
