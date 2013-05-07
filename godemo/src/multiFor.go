package main

//并发循环
//不按照循环体的顺序执行，根据循环中谁先完成，进行返回

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const N = 100

func mfor() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())

	sem := make(chan int, 5)
	chTime := make(chan int, 5)

	//建立协程

	go func() {
		//FOR循环体

		for i := N; i > 0; i-- {
			ii := rand.Intn(10000)

			go func(i, ii int) {

				//doSomething(i, xi)
				time.Sleep(time.Duration(ii) * time.Millisecond)

				//计数
				//vv := i

				sem <- i
				chTime <- ii

			}(i, ii)
		}

		close(sem)
		close(chTime)

	}()

	// 等待循环结束
	//for i := 0; i < N; i++ {
	//	x := <-sem

	//	fmt.Println(x, " ", <-chTime)
	//}

	/**for cc := range sem {
		fmt.Println(cc, " ", <-chTime)
	}*/

	for {
		if cc, ok := <-sem; ok {
			fmt.Println(cc, " ", <-chTime)
		} else {
			//close(sem)
			//close(chTime)
			break
		}
	}

}

func tTimeout() {
	c := make(chan int)
	o := make(chan bool)
	tick := time.Tick(1 * time.Second)
	go func() {
	    t:= time.After(5 * time.Second) // 只⽤用⼀一次，没必要⽤用外部变量。
		for {
			select {
			case v := <-c:
				println(v)
			case <-t:
				println("timeout")
				o <- true
				break
			case  <-tick:
				println("tick")
			}
		}
	}()
	<-o
}

func main() {
	runtime.GOMAXPROCS(2)
	//mfor()

	tTimeout()

}
