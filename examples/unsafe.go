package examples

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func DemoUnsafe() {
	// demoSizeOfInt()
	// demoSizeOfFloat()
	// demoSizeOfString()
	// demoSizeOfSlice()
	demoSizeOfMap()
}

func demoSizeOfInt() {
	var a int8 = math.MaxInt8
	var b int16 = math.MaxInt16
	var c int32 = math.MaxInt32
	var d int64 = math.MaxInt64
	var e uint8 = math.MaxUint8
	var f uint16 = math.MaxUint16
	var g uint32 = math.MaxUint32
	var h uint64 = math.MaxUint64

	fmt.Printf("a(int8): %v bytes\n", unsafe.Sizeof(a))
	fmt.Printf("b(int16): %v bytes\n", unsafe.Sizeof(b))
	fmt.Printf("c(int32): %v bytes\n", unsafe.Sizeof(c))
	fmt.Printf("d(int64): %v bytes\n", unsafe.Sizeof(d))
	fmt.Printf("e(uint8): %v bytes\n", unsafe.Sizeof(e))
	fmt.Printf("f(uint16): %v bytes\n", unsafe.Sizeof(f))
	fmt.Printf("g(uint32): %v bytes\n", unsafe.Sizeof(g))
	fmt.Printf("h(uint64): %v bytes\n", unsafe.Sizeof(h))
}

func demoSizeOfFloat() {
	var a float32 = math.MaxFloat32
	var b float64 = math.MaxFloat64
	fmt.Printf("a(float32): %v bytes\n", unsafe.Sizeof(a))
	fmt.Printf("b(float64): %v bytes\n", unsafe.Sizeof(b))
}

func demoSizeOfString() {
	var s string = "ABC"
	fmt.Printf("unsafe.Sizeof(s): %v\n", unsafe.Sizeof(s))
}

func demoSizeOfSlice() {
	var s []int = []int{1, 2, 3}
	fmt.Printf("s: %v\n", unsafe.Sizeof(s))
}

func demoSizeOfMap() {
	var m map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Printf("s: %v\n", unsafe.Sizeof(m))
	fmt.Printf("t: %v\n", reflect.TypeOf(m).Size())
}
