package examples

import (
	"fmt"
	"strings"
)

func DemoRanges() {
	rangeSlice()
	rangeMap()
	rangeString()
}

func rangeSlice() {
	fmt.Println(newline + fname(rangeSlice))

	s := []int{11, 22, 33, 44, 55}
	r := make([]string, 0, len(s))
	for index := range s {
		r = append(r, fmt.Sprintf("index = %v", index))
	}
	fmt.Println(strings.Join(r, ", "))

	r = make([]string, 0, len(s))
	for index, value := range s {
		r = append(r, fmt.Sprintf("[index = %v, value = %v]", index, value))
	}
	fmt.Println(strings.Join(r, ", "))
}

func rangeMap() {
	fmt.Println(newline + fname(rangeMap))
	m := map[string]any{
		"name":    "Kirill",
		"age":     26,
		"married": false,
	}

	r := make([]string, 0, len(m))
	for key := range m {
		r = append(r, fmt.Sprintf("%v = %v", key, m[key]))
	}
	fmt.Println(strings.Join(r, ", "))

	r = make([]string, 0, len(m))
	for key, value := range m {
		r = append(r, fmt.Sprintf("%v = %v", key, value))
	}
	fmt.Println(strings.Join(r, ", "))

}

func rangeString() {
	fmt.Println(newline + fname(rangeString))

	const s = "ABCDEFG"

	for index := range s {
		fmt.Println("index =", index)
	}

	r := make([]string, 0, len(s))
	for index, value := range s {
		r = append(r, fmt.Sprintf("index = %v, value = %v, string(value) = %v", index, value, string(value)))
	}

	for _, v := range r {
		fmt.Println(v)
	}
}
