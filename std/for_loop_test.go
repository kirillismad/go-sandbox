package std

import (
	"context"
	"iter"
	"testing"
	"time"

	"github.com/samber/lo"
)

func TestLoop(t *testing.T) {
	t.Parallel()
	t.Run("ClassicForLoop", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < 10; i++ {
			t.Logf("i = %d", i)
		}

	})
	t.Run("WhileLoop", func(t *testing.T) {
		t.Parallel()

		x := 100
		for x >= 0 {
			x = x - getRandomNumber()
			t.Logf("x = %d", x)
		}
	})

	t.Run("DoWhileLoop", func(t *testing.T) {
		t.Parallel()

		x := 100
		for {
			t.Logf("x = %d", x)
			x = x - getRandomNumber()

			if x < 0 {
				break
			}
		}
	})

	t.Run("InfiniteLoop", func(t *testing.T) {
		t.Parallel()

		start := time.Now()
		timeout := 2 * time.Second

		for {
			if time.Since(start) > timeout {
				break
			}
			t.Logf("rand = %d", getRandomNumber())
			time.Sleep(500 * time.Millisecond)
		}
	})
	t.Run("RangeOverInt", func(t *testing.T) {
		t.Parallel()

		for i := range 10 {
			t.Logf("i = %d", i)
		}
	})
}

func TestRangeOverString(t *testing.T) {
	t.Parallel()

	asciiStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	unicodeStr := "Господа, базовый вектор развития требует анализа новых принципов формирования материально-технической и кадровой базы."

	t.Run("Ascii", func(t *testing.T) {
		t.Parallel()

		for i, r := range asciiStr {
			t.Logf("i = %d, r = %c", i, r)
		}
	})

	t.Run("Unicode", func(t *testing.T) {
		t.Parallel()

		for i, r := range unicodeStr {
			t.Logf("i = %d, r = %c", i, r)
		}
	})
}

func TestRangeOverSlice(t *testing.T) {
	t.Parallel()

	slice := []int{5, 4, 6, 7, 8, 3}
	t.Run("Default", func(t *testing.T) {
		t.Parallel()

		for i, v := range slice {
			t.Logf("i = %d, v = %d", i, v)
		}
	})
	t.Run("Value", func(t *testing.T) {
		t.Parallel()

		for _, v := range slice {
			t.Logf("v = %d", v)
		}
	})
	t.Run("Index", func(t *testing.T) {
		t.Parallel()

		for i := range slice {
			t.Logf("i = %d", i)
		}
	})

	t.Run("Nil", func(t *testing.T) {
		t.Parallel()

		var nilSlice []int

		for i, v := range nilSlice {
			t.Log(i, v)
		}
	})
}

func TestRangeOverMap(t *testing.T) {
	t.Parallel()

	t.Run("KeyValue", func(t *testing.T) {
		t.Parallel()

		for key, value := range constantsExample() {
			t.Logf("%s = %f", key, value)
		}
	})
	t.Run("Value", func(t *testing.T) {
		t.Parallel()

		for _, value := range constantsExample() {
			t.Logf("value = %f", value)
		}
	})
	t.Run("Key", func(t *testing.T) {
		t.Parallel()

		for key := range constantsExample() {
			t.Logf("key = %s", key)
		}
	})
	t.Run("Nil", func(t *testing.T) {
		t.Parallel()

		var nilMap map[string]int
		for k, v := range nilMap {
			t.Log(k, v)
		}
	})
}

func TestRangeOverChannel(t *testing.T) {
	t.Parallel()
	t.Run("Default", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := range 3 {
				ch <- i
			}
		}()

		for i := range ch {
			t.Logf("i = %d", i)
		}
	})
	t.Run("EmptyChannel", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		inner := func() struct{} {
			ch := make(chan int)
			for v := range ch {
				t.Log(v)
			}
			return struct{}{}
		}

		select {
		case <-ctx.Done():
			t.Log("Context done")
		case <-lo.Async(inner):
			t.Log("Never executed")
		}

	})
}

func TestRangeOverFunc(t *testing.T) {
	t.Parallel()

	constPairs := func() iter.Seq2[string, float64] {
		return func(yield func(string, float64) bool) {
			for k, v := range constantsExample() {
				if !yield(k, v) {
					return
				}
			}
		}
	}

	fib := func(n int) iter.Seq[int] {
		return func(yield func(int) bool) {
			x1, x2 := 0, 1
			for i := 0; i < n; i++ {
				if !yield(x1) {
					return
				}
				x1, x2 = x2, x1+x2
			}
		}
	}

	t.Run("iter.Seq", func(t *testing.T) {
		t.Parallel()

		for x := range fib(10) {
			t.Logf("x = %d", x)
		}
	})
	t.Run("iter.Seq(break)", func(t *testing.T) {
		t.Parallel()

		cnt := 0
		for x := range fib(10) {
			if cnt == 3 {
				break
			}
			t.Logf("x = %d", x)
			cnt++
		}
	})
	t.Run("iter.Seq2", func(t *testing.T) {
		t.Parallel()

		for k, v := range constPairs() {
			t.Logf("k = %s, v = %f", k, v)
		}
	})
	t.Run("iter.Pull", func(t *testing.T) {
		t.Parallel()

		iterator := fib(10)
		next, stop := iter.Pull(iterator)
		defer stop()

		for {
			v, ok := next()
			if !ok {
				break
			}
			t.Logf("v = %d", v)
		}

	})
	t.Run("iter.Pull2", func(t *testing.T) {
		t.Parallel()

		iterator := constPairs()
		next, stop := iter.Pull2(iterator)
		defer stop()

		for {
			k, v, ok := next()
			if !ok {
				break
			}
			t.Logf("k = %s, v = %f", k, v)
		}
	})
}
