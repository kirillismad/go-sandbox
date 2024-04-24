package examples

import (
	"fmt"
	"strings"
)

func DemoMaps() {
	// nilMap()
	// initMap()
	missingKey()
	// setMap()
	// deleteMap()
	// iterMap()
	// extendMap()
}

func initMap() {
	fmt.Println(newline + fname(initMap))
	var m1 map[string]int
	fmt.Printf("m1 = %v, len(m1) = %v, m1 == nil = %v\n", m1, len(m1), m1 == nil)
	// m1["key"] = 42 // panic
	v, ok := m1["key"]
	fmt.Println("v:", v, "ok:", ok)

	m2 := map[string]int{}
	fmt.Printf("m2 = %v, len(m2) = %v, m2 == nil = %v\n", m2, len(m2), m2 == nil)

	m3 := make(map[string]int)
	fmt.Printf("m3 = %v, len(m3) = %v, m3 == nil = %v\n", m3, len(m3), m3 == nil)

	m4 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Printf("m4 = %v, len(m4) = %v, m4 == nil = %v\n", m4, len(m4), m4 == nil)

}

func nilMap() {
	var m map[string]int
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func missingKey() {
	fmt.Println(newline + fname(missingKey))

	var m map[string]int
	var value int
	var isPresent bool

	value = m["key"] // no error
	fmt.Println("value =", value)

	value, isPresent = m["key"] // no error
	fmt.Println("value =", value, "isPresent =", isPresent)
}

func setMap() {
	fmt.Println(newline + fname(setMap))

	m := make(map[string]int)
	m["key"] = 1
	fmt.Printf("m = %v, len(m) = %v, m == nil = %v\n", m, len(m), m == nil)
}

func deleteMap() {
	fmt.Println(newline + fname(deleteMap))
	m := map[string]int{"key": 1}
	delete(m, "key")
	fmt.Printf("m = %v, len(m) = %v, m == nil = %v\n", m, len(m), m == nil)
}

func iterMap() {
	fmt.Println(newline + fname(iterMap))
	m := map[string]any{
		"name":    "Kirill",
		"age":     26,
		"married": false,
	}
	s := make([]string, 0, len(m))
	for key := range m {
		s = append(s, fmt.Sprintf("%v = %v", key, m[key]))
	}
	fmt.Println(strings.Join(s, ", "))

}

func extendMap() {
	fmt.Println(newline + fname(extendMap))

	m1 := map[string]any{
		"name":    "Kirill",
		"age":     26,
		"married": false,
	}
	m2 := map[string]any{
		"surname": "Pech",
		"salary":  "61 nzd/h",
	}

	for k := range m2 {
		m1[k] = m2[k]
	}

	s := make([]string, 0, len(m1))
	for k, v := range m1 {
		s = append(s, fmt.Sprintf("%v = %v", k, v))
	}
	fmt.Println(strings.Join(s, ", "))
}
