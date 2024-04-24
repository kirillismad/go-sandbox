package examples

import (
	"fmt"
	"unicode/utf8"
)

func DemoStrings() {
	stringLen()
	stringTypes()
}

func stringLen() {
	fmt.Println(newline + fname(stringLen))

	const s1 = "Hello"
	fmt.Println("len(s1) =", len(s1))
	fmt.Println("utf8.RuneCountInString(s1) =", utf8.RuneCountInString(s1))
	for i, v := range s1 {
		fmt.Printf("i = %v, v = %v, string(v) = %v \n", i, v, string(v))
	}

	const s2 = "АБВГ"
	fmt.Println("len(s2) =", len(s2))
	fmt.Println("utf8.RuneCountInString(s2) =", utf8.RuneCountInString(s2))
	for i, v := range s2 {
		fmt.Printf("i = %v, v = %v, string(v) = %v \n", i, v, string(v))
	}
}

func stringTypes() {
	s1 := "doble\tquote"
	s2 := `strange\tquote`
	s3 := 'r'
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
