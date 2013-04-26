package main

import (
	"fmt"
)

const N = 100

func mfor() {
	sem := make(chan int, N)

	//FOR循环体

	for i, xi := range data {

		//建立协程

		go func(i int, xi float) {

			//doSomething(i, xi)

			//计数

			sem <- 0

		}(i, xi)

	}

	// 等待循环结束
	for i := 0; i < N; i++ {
		x <- sem
	}

}

func main() {

}
