package examples

import (
	"fmt"
	"time"
)

func DemoChannels() {
	// demoUnbufChannel()
	// demoBufChannel()
	demoChannelRead()
}

func sender1(channel chan int) {
	fmt.Println("Sender: before")
	channel <- rint(100)
	fmt.Println("Sender: after")
}

func receiver1(channel <-chan int, result chan int) {
	fmt.Println("Receiver: before")
	<-channel
	fmt.Println("Receiver: after")
	result <- 0
}

func demoUnbufChannel() {
	channel := make(chan int)
	result := make(chan int)
	go sender1(channel)
	go receiver1(channel, result)

	<-result
}

func sender2(channel chan<- int) {
	const ITERS = 3
	for i := 0; i < ITERS; i++ {
		channel <- rint(100)
	}
	close(channel)
	fmt.Println("Closed")
}

func empty() struct{} {
	return struct{}{}
}

func receiver2(channel <-chan int, result chan<- struct{}) {
	time.Sleep(1 * time.Second)
	for e := range channel {
		fmt.Println(e)
	}
	result <- empty()
}

func demoBufChannel() {
	channel := make(chan int, 3)
	result := make(chan struct{})
	go sender2(channel)
	go receiver2(channel, result)
	<-result
}

func demoChannelRead() {
	ch := make(chan int)

	x, ok := <-ch
	fmt.Printf("x: %v, ok: %v\n", x, ok)
}
