package examples

import (
	"fmt"
	"math"
)

func DemoPointers() {
	// demoInt()
	// demoSlice()
	// demoMaps()
	// demoPointerToPointer()
	demoIntPointerForLoop()
}

func intFuncValue(value int) int {
	fmt.Printf("intFuncValue: %p\n", &value)
	value = 111
	return value
}

func intFuncPointer(pointer *int) *int {
	fmt.Printf("intFuncPointer: %p\n", pointer)
	*pointer = 222
	return pointer
}

func demoInt() {
	v1, v2 := 11, 22
	fmt.Printf("Before: &v1 = %p, &v2 = %p\n", &v1, &v2)
	v3 := intFuncValue(v1)
	v4 := intFuncPointer(&v2)
	fmt.Printf("After: &v1 = %p, &v2 = %p\n", &v1, &v2)
	fmt.Printf("v3 = %p, &v4 = %p\n", &v3, v4)
}

func sliceFunc[T any](s []T) []T {
	s = append(s, s[len(s)-1])
	return s
}
func demoSlice() {
	s := make([]int, 5)
	for i := range s {
		s[i] = i
	}
	retS := sliceFunc(s)
	s[0] = 11

	fmt.Printf("&s = %p, s = %v, &retS = %p, retS = %v\n", &s, s, &retS, retS)

}

func mapFunc(m map[string]int) map[string]int {
	m["THREE"] = 33
	m["FOUR"] = 4
	return m
}
func min(nums ...int) int {
	min := math.Inf(1)
	for _, num := range nums {
		min = math.Min(float64(num), min)
	}
	return int(min)
}
func demoMaps() {
	keys := []string{"ONE", "TWO", "THREE"}
	values := []int{1, 2, 3, 4}

	m := make(map[string]int, min(len(keys), len(values)))

	for i := 0; i < min(len(keys), len(values)); i++ {
		m[keys[i]] = values[i]
	}

	retM := mapFunc(m)
	fmt.Printf("&m = %p, m = %v, &retM = %p, retM = %v\n", &m, m, &retM, retM)
}

func demoPointerToPointer() {
	x := 10
	ptr_x := &x
	ptr_ptr_x := &ptr_x

	fmt.Println(x, ptr_x, ptr_ptr_x)
	fmt.Println(**ptr_ptr_x)
}

func demoIntPointerForLoop() {
	var l1 []*int

	for _, v := range []int{1, 2, 3, 4, 5} {
		vv := v // new address
		l1 = append(l1, &vv)
	}

	for _, v := range l1 {
		fmt.Printf("%v, %p ", *v, v)
	}

}
