package tasks

import "fmt"

func task1() {
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			x := <-numbers
			squares <- x * x
		}
		close(squares)
	}()

	for e := range squares {
		fmt.Println(e)
	}
	fmt.Println("END")
}
