package examples

import (
	"fmt"
	"math"
	"strconv"
)

func DemoBitwise() {
	fmt.Println(math.MaxInt8)
	fmt.Println(strconv.FormatInt(int64(1<<7-1), 2))
}
