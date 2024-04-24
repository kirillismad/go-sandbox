package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func DemoSelect() {
	// demoSelect()
	demoSelectFromClosedChannel()
}

func demoSelect() {
	var channels []chan string

	const producers = 5

	done := make(chan int)

	for i := 0; i < producers; i++ {
		ch := make(chan string)
		channels = append(channels, ch)

		go func(i int) {
			logProducer(i, ch)
			done <- i
		}(i)
	}
	consumerDone := make(chan bool)
	go logsConsumer(channels, consumerDone)

	for i := 0; i < producers; i++ {
		id := <-done
		fmt.Println("Producer", id, "finished.")
	}
	consumerDone <- true
}

func logProducer(id int, ch chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(1000+1)) * time.Millisecond)
		ch <- fmt.Sprintf("Producer: %v, i: %v", id, i)
	}
}

func logsConsumer(channels []chan string, done chan bool) {
	for {
		for i := range channels {
			select {
			case s := <-channels[i]:
				fmt.Println(s)
			case <-done:
				return
			default:

			}
		}
	}
}

func demoSelectFromClosedChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	done := make(chan struct{})

	for i := 0; i < 4; i++ {
		select {
		case <-done:
			return
		default:
			x, ok := <-ch
			fmt.Printf("x: %v ok: %v\n", x, ok)
		}
	}

}
