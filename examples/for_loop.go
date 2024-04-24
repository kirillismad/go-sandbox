package examples

import (
	"cmp"
	"fmt"
	"math/rand"
)

func DemoLoop() {
	const count = 10
	forLoop(count)
	whileLoop(count)
	infLoop(count)
	fibonacciLoop()
	forOverString()
	demoFionacciLoopChan()
	demoQuickSearch()
	demoQuickSort()
	demoBubbleSort()
	demoMin1()
	demoTwoPointersLoop1()
	demoTwoPointersLoop2()
	demoForLoopOverStringWithIndex()
	demoReverse()
}

func forLoop(value int) {
	fmt.Println("\nForLoop")
	for i := 0; i < value; i++ {
		fmt.Println("i =", i)
		if i%2 == 0 {
			continue
		}
	}
}

func whileLoop(value int) {
	fmt.Println("\nWhile loop")
	x := 0
	for x < value {
		fmt.Println("x =", x)
		x++
	}
}

func infLoop(value int) {
	fmt.Println("\nInfLoop")
	x := 0

	for {
		if x == value {
			break
		}
		fmt.Println("x =", x)
		x++
	}
}

func forOverString() {
	fmt.Println("\nforOverString")
	str := "Строка на русском языке"

	for z, x := range str {
		fmt.Println(z, x, string(x))
	}
}

func fibonacciLoop() {
	fmt.Println(newline + fname(fibonacciLoop))
	const count = 10
	fmt.Printf("Print first %v fibonacci numbers\n", count)

	i1, i2 := 0, 1
	var next int
	for i := 0; i < count; i++ {
		next, i1, i2 = i1, i2, i1+i2
		fmt.Printf("%v) %v\n", i+1, next)
	}
}

func fibonacciLoopChan(count int, channel chan<- int) {
	next, n1, n2 := 0, 0, 1
	for i := 0; i < count; i++ {
		next, n1, n2 = n1, n2, n1+n2
		channel <- next
	}
	close(channel)
}

func demoFionacciLoopChan() {
	const count = 10
	channel := make(chan int)
	go fibonacciLoopChan(count, channel)
	i := 1
	for e := range channel {
		fmt.Printf("%v) %v\n", i, e)
		i++
	}
}

// [1, 2, 3, 4, 5, 6, 7], t = 7
// l = 0, r = 6, m = 3
// l = 4, r = 6, m = 5
// l = 6, r = 6, m = 6
func quickSearch(slice []int, target int) int {
	left := 0
	right := len(slice) - 1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println("\t", "left:", left, "mid:", mid, "right:", right)
		if target == slice[mid] {
			return mid
		} else if target < slice[mid] {
			right = mid - 1
		} else if target > slice[mid] {
			left = mid + 1
		}
	}
	return -1
}

func demoQuickSearch() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	cases := []struct {
		slice  []int
		target int
		result int
	}{
		{slice, 1, 0},
		{slice, 2, 1},
		{slice, 3, 2},
		{slice, 4, 3},
		{slice, 5, 4},
		{slice, 6, 5},
		{slice, 7, 6},
	}

	for i, v := range cases {
		r := quickSearch(v.slice, v.target)
		fmt.Println(i, ")", "slice:", v.slice, "target:", v.target, "result:", v.result, "r:", r, v.result == r)
	}
}

func partition(slice []int, left int, right int) int {
	pivot := slice[(left+right)/2]
	for {
		for slice[left] < pivot {
			left++
		}
		for slice[right] > pivot {
			right--
		}
		if left >= right {
			return right
		}
		slice[left], slice[right] = slice[right], slice[left]
	}
}

// [6, 4, 7, 0, 1, 2, 3, 9, 5, 8]
func quickSort(slice []int, left int, right int) {
	if left < right {
		p := partition(slice, left, right)
		quickSort(slice, left, p)
		quickSort(slice, p+1, right)
	}
}

func demoQuickSort() {
	slice := rand.Perm(5)
	fmt.Println(slice)
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func demoBubbleSort() {
	slice := rand.Perm(10)
	fmt.Println(slice)
	bubbleSort(slice)
	fmt.Println(slice)
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
	}
}

func min1[T cmp.Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("error")
	}

	m := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < m {
			m = slice[i]
		}
	}
	return m
}

func demoMin1() {
	s := make([]int, 5)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
	fmt.Println(s)
	m := min1(s)
	fmt.Println(m)
}

func demoTwoPointersLoop1() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		fmt.Printf("i=%v, j=%v, s[i]=%v, s[j]=%v\n", i, j, s[i], s[j])
	}
}

func demoTwoPointersLoop2() {
	fmt.Println()
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1
		fmt.Printf("i=%v, j=%v, s[i]=%v, s[j]=%v\n", i, j, s[i], s[j])
	}
}

func demoForLoopOverStringWithIndex() {
	fmt.Println()
	s := "QЭRЮSЯ"
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		j := len(r) - i - 1
		fmt.Printf("i=%v, j=%v, s[i]=%c, s[j]=%c\n", i, j, r[i], r[j])
	}
}

func demoReverse() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(s)
	fmt.Println(s)
}

func reverse[T any](s []T) {
	totalLen := len(s)
	for i, j := 0, totalLen-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
