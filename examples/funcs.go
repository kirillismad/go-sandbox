package examples

import "fmt"

func DemoFuncs() {
	// typedArg(rint(100))
	// sameTypeArgs(rint(100), rint(100))
	// r1 := simpleReturn()
	// r2, r3 := multiReturn()
	// r4 := namedReturn()
	// variadicFunc([]int{1, 2, 3, 4, 5}...)

	// fmt.Println()
	// fmt.Println("retValue() =", r1)
	// fmt.Println("multiReturn() =", r2, r3)
	// fmt.Println("namedReturn() =", r4)
	// funcAsVar()
	demoDirtyVariadic()
}

func typedArg(arg int) {
	fmt.Println(newline + fname(typedArg))
	fmt.Println("arg =", arg)
}

func sameTypeArgs(arg1, arg2 int) {
	fmt.Println(newline + fname(sameTypeArgs))
	fmt.Printf("arg1 = %v, arg2 = %v\n", arg1, arg2)
}

func simpleReturn() int {
	fmt.Println(newline + fname(simpleReturn))
	ret := rint(100)
	fmt.Printf("ret = %v\n", ret)
	return ret
}

func multiReturn() (string, int) {
	fmt.Println(newline + fname(multiReturn))
	ret1 := "abc"
	ret2 := rint(100)
	fmt.Printf("ret1 = %v, ret2 = %v\n", ret1, ret2)
	return ret1, ret2
}

func namedReturn() (n int) {
	fmt.Println(newline + fname(namedReturn))
	n = rint(100)
	fmt.Println("n =", n)
	return
}

func variadicFunc(nums ...int) {
	fmt.Println(newline + fname(variadicFunc))
	for i, v := range nums {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}
}

func funcAsVar() {
	fmt.Println(newline + fname(funcAsVar))
	var intSum func(int, int) int = opSum[int]
	intSub := opSub[int]

	intOps := map[rune]func(int, int) int{
		'+': intSum,
		'-': intSub,
	}

	floatOps := map[rune]func(float64, float64) float64{
		'+': opSum[float64],
		'-': opSub[float64],
	}

	intA, intB := 6, 2
	floatA, floatB := 7.34, 3.21

	for _, opSign := range []rune{'+', '-'} {
		intOp := intOps[opSign]
		floatOp := floatOps[opSign]
		fmt.Println(intA, string(opSign), intB, "=", intOp(intA, intB))
		fmt.Println(floatA, string(opSign), floatB, "=", floatOp(floatA, floatB))
	}
}

func opSum[T int | float64](a, b T) T {
	return a + b
}

func opSub[T int | float64](a, b T) T {
	return a - b
}

func dirtyVariadic(s ...int) {
	s[0] = 42
}

func demoDirtyVariadic() {
	s := []int{1, 2, 3}
	dirtyVariadic(s...)
	fmt.Printf("s: %v\n", s)
}
