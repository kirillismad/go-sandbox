package std

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaps(t *testing.T) {
	t.Parallel()
	t.Run("comparable types: chan", func(t *testing.T) {
		t.Parallel()
		r := require.New(t)

		ch1 := make(chan int)
		ch2 := make(chan int)

		chanMap := map[chan int]int{
			ch1: 1,
			ch2: 2,
		}

		for k, v := range chanMap {
			go func() {
				t.Logf("%d", v)
				k <- v
			}()
		}

		for k := range chanMap {
			v := <-k
			r.Equal(chanMap[k], v)
		}
	})
}
