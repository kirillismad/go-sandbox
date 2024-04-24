package examples

import "fmt"

func DemoRecursion() {
	demoFubonacciRecursion()
}

// 0 1 1 2 3 5 8 13
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func demoFubonacciRecursion() {
	fmt.Println(newline + fname(demoFubonacciRecursion))
	const count = 10
	fmt.Printf("Print first %v fibonacci numbers\n", count)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v) %v\n", i+1, fibonacci(i))
	}
}
