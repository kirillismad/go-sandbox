package examples

import "fmt"

func DemoComparison() {
	demoInterfaces()
}

type xInterface interface {
	doX()
}

type xDoerInt int

func (x xDoerInt) doX() {
	fmt.Println(x)
}

type xDoerString string

func (x xDoerString) doX() {
	fmt.Println(x)
}

func compareXInterfaces(x1, x2 xInterface) bool {
	return x1 == x2
}

func demoInterfaces() {
	x1 := xDoerInt(42)
	x11 := xDoerInt(43)
	x2 := xDoerString("string1")
	x22 := xDoerString("string2")

	result := compareXInterfaces(x1, x2)
	fmt.Println(result)

	result = compareXInterfaces(x1, x11)
	fmt.Println(result)

	result = compareXInterfaces(x2, x22)
	fmt.Println(result)
}
