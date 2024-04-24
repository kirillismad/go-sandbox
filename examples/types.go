package examples

import "fmt"

type Int1 int
type Int11 Int1

type Int2 int
type Int22 Int2

func DemoTypes() {
	var i int = 1

	var i1 Int1 = Int1(i)
	var i11 Int11 = Int11(i1)

	var i2 Int2 = Int2(i)
	var i22 Int22 = Int22(i2)

	fmt.Println(Int22(i11) == i22)
	fmt.Println(Int11(i22) == i11)
}
