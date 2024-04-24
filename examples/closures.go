package examples

import "fmt"

func DemoClosures() {
	fibonacciDemo()
}

func fibonacciDemo() {
	fmt.Println(newline + fname(fibonacciDemo))
	const count = 10
	fmt.Printf("Print first %v fibonacci numbers\n", count)
	var next func() int = fibonacciFactory()
	for i := 0; i < count; i++ {
		fmt.Printf("%v) %v\n", i+1, next())
	}
}

// 0 1 1 2 3 5 8 13
func fibonacciFactory() func() int {
	i1 := 0
	i2 := 1

	var next int
	return func() int {
		next, i1, i2 = i1, i2, i1+i2
		return next
	}
}
