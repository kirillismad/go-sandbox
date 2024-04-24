package examples

import (
	"fmt"
	"time"
)

func DemoSwitch() {
	simpleSwitch()
	binarySwitch()
	multiSwitch()
	typeSwitch(42)
	typeSwitch(42.32)
	typeSwitch("Hello world")
	typeSwitch(true)
}

func simpleSwitch() {
	fmt.Println("\nsimpleSwitch")
	switch x := rint(3); x {
	case 1: // x == 1
		fmt.Println("RED")
	case 2: // x == 2
		fmt.Println("YELLOW")
	case 3: // x == 3
		fmt.Println("GREEN")
	default: // else
		fmt.Println("INVALID")
	}
}

func binarySwitch() {
	fmt.Println("\nbinarySwitch")
	x := rint(100)
	switch { //same as "switch true"
	case x <= 50:
		fmt.Println("FIRST HALF")
	case x > 50:
		fmt.Println("SECOND HALF")
	}
}

func multiSwitch() {
	fmt.Println("\nmultiSwitch")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // value1, value2, ..., valueN
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a working day")
	}
}

func typeSwitch(value any) {
	fmt.Println("\ntypeSwitch")
	switch value.(type) {
	case bool:
		fmt.Println("BOOL")
	case int, float64:
		fmt.Println("INT OR FLOAT")
	case string:
		fmt.Println("STRING")
	default:
		fmt.Println("UNKNOWN")
	}
}
