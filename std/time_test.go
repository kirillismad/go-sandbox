package std

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Parallel()

	maxSpeed := 10.0 / float64(time.Second.Nanoseconds())

	t2 := 10 * time.Second

	s := maxSpeed * float64(t2.Nanoseconds())

	t.Log(s)
}
