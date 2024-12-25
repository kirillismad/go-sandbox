package std

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

/*
Отправка (ch <- x) "synchronized-before" соответствующего получения (<-ch), так как отправка завершится только тогда, когда кто-то прочитает значение из канала.
Получение (<-ch) "synchronized-before" продолжения выполнения отправляющей горутины после завершения получения.

Чтение:
- из небуф канала блокируется, пока не будет записано значение
- из буф канала блокируется, если буфер пуст, иначе возвращает значение и true
- из nil канала блокируется навсегда
- из закрытого канала возвращает нулевое значение и false
- из закрытого буферизованного канала возвращает значение и true, если буфер не пуст, иначе нулевое значение и false

Запись:
- в небуф канал блокируется, пока не будет прочитано значение
- в буф канал блокируется, если буфер полон, иначе записывает значение
- в nil канал блокируется навсегда
- в закрытый (буф/небуф) канал вызывает панику

Закрытие:
- nil канал вызывает панику
- закрытый канал вызывает панику
*/

func TestChannels(t *testing.T) {
	t.Parallel()

	// последовательность
	// 1. G1. отправка в канал
	// 2. G2. получение из канала
	// 3. G2. got = x
	// 4. G1. продолжение выполнения
	t.Run("synchronized before", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)
		t.Cleanup(func() {
			close(ch)
		})
		x := getRandomNumber()
		go func() {
			// Отправка (ch <- x) "synchronized-before" соответствующего получения (<-ch), так как отправка завершится только тогда, когда кто-то прочитает значение из канала.
			ch <- x
		}()

		time.Sleep(1 * time.Second)
		// Получение (<-ch) "synchronized-before" продолжения выполнения отправляющей горутины после завершения получения.
		got := <-ch
		require.Equal(t, x, got)
	})

	t.Run("read, unbuf", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)

		t.Cleanup(func() {
			close(ch)
		})

		select {
		// block until write
		case val, ok := <-ch:
			t.Fatalf("unexpected, %d, %v", val, ok)
		default:
		}
	})
	t.Run("read, buf", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int, 1)
		t.Cleanup(func() {
			close(ch)
		})

		// block until write
		select {
		case val, ok := <-ch:
			t.Fatalf("unexpected, %d, %v", val, ok)
		default:
		}
	})
	t.Run("read, nil", func(t *testing.T) {
		t.Parallel()

		var ch chan int

		select {
		// block forever
		case val, ok := <-ch:
			t.Fatalf("unexpected, %d, %v", val, ok)
		default:
		}
	})
	t.Run("read, closed", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)
		close(ch)

		val, ok := <-ch
		require.Zero(t, val)
		require.False(t, ok)
	})

	t.Run("read from buffered closed", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int, 1)
		ch <- getRandomNumber()
		close(ch)

		val, ok := <-ch
		require.NotZero(t, 1, val)
		require.True(t, ok)

		val, ok = <-ch
		require.Zero(t, val)
		require.False(t, ok)
	})

	t.Run("write, unbuf", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)

		t.Cleanup(func() {
			close(ch)
		})

		// block until read
		select {
		case ch <- getRandomNumber():
			t.Fatalf("unexpected")
		default:
		}
	})
	t.Run("write, buf", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int, 1)
		t.Cleanup(func() {
			close(ch)
		})

		// not blocking if buffer is not full
		ch <- getRandomNumber()

		// block until read
		select {
		case ch <- getRandomNumber():
			t.Fatalf("unexpected")
		default:
		}
	})

	t.Run("write, nil", func(t *testing.T) {
		t.Parallel()

		var ch chan int

		// block forever
		select {
		case ch <- getRandomNumber():
			t.Fatalf("unexpected")
		default:
		}
	})
	t.Run("write, closed no matter buf or unbuf", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)
		close(ch)

		// panic
		require.Panics(t, func() {
			ch <- getRandomNumber()
		})
	})

	t.Run("close, nil", func(t *testing.T) {
		t.Parallel()

		var ch chan int

		// panic
		require.Panics(t, func() {
			close(ch)
		})
	})
	t.Run("close, closed", func(t *testing.T) {
		t.Parallel()

		ch := make(chan int)
		close(ch)

		// panic
		require.Panics(t, func() {
			close(ch)
		})
	})
}
