package examples

import "fmt"

func DemoArrays() {
	initArray()
	arrayIter()
	twoDimensions()
	threeDimensions()
}

func initArray() {
	fmt.Println("\ninitArray")
	var d [3]int
	fmt.Println("d =", d)
	d[0] = 42
	fmt.Println("d =", d)
	fmt.Println("d[0] =", d[0])

	a := [3]int{1, 2} // [1, 2, 0]
	fmt.Println("a =", a)
	a = [3]int{3, 2, 1}
	fmt.Println("a =", a)

	arr := [...]int{1, 2, 3}
	fmt.Println("arr =", arr)
}

func arrayIter() {
	fmt.Println("\narrayIter")
	arr := [3]int{11, 22, 33}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%v] = %v\n", i, arr[i])
	}
}

func twoDimensions() {
	fmt.Println("\ntwoDimensions")
	const d1 = 2
	const d2 = 3
	var arr [d1][d2]int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = j + i*len(arr[i])
		}
	}
	fmt.Println("arr =", arr)
}

func threeDimensions() {
	fmt.Println("\nthreeDimensions")
	const d = 2
	var arr [d][d][d]int
	x := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			for k := 0; k < len(arr[i][j]); k++ {
				arr[i][j][k] = x
				x++
			}
		}
	}
	fmt.Println(arr)
}
