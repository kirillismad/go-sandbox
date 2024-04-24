package examples

import (
	"fmt"
)

func DemoIfElse() {
	shortIf(-1)
	commonIf(rint(100))
	ifElseIf("GREEN")
	ifAndInit()
	ifRuneAndInt32()
}

func shortIf(value int) {
	fmt.Println("\nshortIf")
	if value < 0 {
		fmt.Println(value, "is unacceptable")
	}
}

func commonIf(value int) {
	fmt.Println("\ncommonIf")
	if value%2 == 0 {
		fmt.Println("value =", value, "(even)")
	} else {
		fmt.Println("value =", value, "(odd)")
	}
}

func ifElseIf(value string) {
	fmt.Println("\nifElseIf")
	if value == "RED" {
		fmt.Println("STOP")
	} else if value == "YELLOW" {
		fmt.Println("READY")
	} else if value == "GREEN" {
		fmt.Println("GO")
	} else {
		fmt.Println("Invalid color:", value)
	}
}

func ifAndInit() {
	fmt.Println("\nifAndInit")
	if x := rint(100); x <= 50 {
		fmt.Println("FIRST HALF")
	} else {
		fmt.Println("SECOND HALF")
	}
}

func ifRuneAndInt32() {
	fmt.Println("\nifRuneAndInt32")
	var x1 rune = 'a'
	var x2 int32 = 97
	if x1 == x2 {
		fmt.Println("Hello world")
	}
}
