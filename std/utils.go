package std

import (
	"math"
	"math/rand"
)

func getRandomNumber() int {
	return rand.Intn(9) + 1
}

func constantsExample() map[string]float64 {
	return map[string]float64{
		"Pi":  math.Pi,
		"Phi": math.Phi,
		"E":   math.E,
	}
}
