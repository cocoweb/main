package main

//并发循环
//不按照循环体的顺序执行，根据循环中谁先完成，进行返回

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100

func mfor() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())

	sem := make(chan int, N)
	chTime := make(chan int)

	//FOR循环体

	for i := N; i > 0; i-- {

		//建立协程

		go func(i int) {
			ii := rand.Intn(10000)

			//doSomething(i, xi)
			time.Sleep(time.Duration(ii) * time.Millisecond)

			vv := i

			//计数

			sem <- vv
			chTime <- ii

		}(i)

	}

	// 等待循环结束
	for i := 0; i < N; i++ {
		x := <-sem

		fmt.Println(x, " ", <-chTime)
	}

}

func main() {
	mfor()

}
