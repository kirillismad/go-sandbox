package examples

import "fmt"

const (
	one = iota
	two
	three
	four
)

const (
	_1  = 1 << iota // 1 << 0 = 1
	_2              // 1 << 1 = 10
	_4              // 1 << 2 = 100
	_8              // 1 << 3 = 1000
	_16             // 1 << 4 = 10000
)

func DemoIota() {
	fmt.Println(one, two, three, four)
	fmt.Println(_1, _2, _4, _8, _16)
}
