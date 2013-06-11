package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		x, y = y, x+y
		select {
		case c <- x:

			fmt.Println("here")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
