package std

import "math/rand"

func getRandomNumber() int {
	return rand.Intn(9) + 1
}
