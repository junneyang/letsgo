package main

import (
	"fmt"
	"runtime"
	//	"sync"
	"time"
)

//var wg sync.WaitGroup // 1

func sayHello(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str, "--", i)
		time.Sleep(1 * time.Second)
	}
}

func sayHi(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str, "--", i)
		//		runtime.Gosched()
	}
}

func sayCh(str string, ch chan int) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str, "--", i)
		//		runtime.Gosched()
	}
	ch <- 888
}

func sum(a []int, c chan int) {
	//	defer wg.Done() // 3

	total := 0
	for _, v := range a {
		total += v
	}
	runtime.Gosched()
	c <- total // send total to c
	//	close(c)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci_new(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("STUCK, NO SIGNAL")
		}
	}
}

func main() {
	fmt.Println("PROCS_NUM: ", runtime.NumCPU())
	fmt.Println("PRE_GOMAXPROCS: ", runtime.GOMAXPROCS(12))
	fmt.Println("CUR_GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	//	go sayHello("hello")
	//	sayHello("world")
	//	go sayHi("hi")
	//	sayHi("why")

	//	var ch chan int = make(chan int)
	//	fmt.Printf("%#v\n", ch)
	//	go sayCh("junneyang", ch)
	//	fmt.Println(<-ch)

	a := []int{7, 2, 8, -9, 4, 0, 6, 74, 12, 88}
	c := make(chan int)
	//	wg.Add(4)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	//	x, y := <-c, <-c // receive from c
	x := <-c
	y := <-c
	m := <-c
	n := <-c
	fmt.Println(x, y, m, n)
	//	for ii := range c { // until channel c closed
	//		fmt.Println(ii)
	//	}
	//	wg.Wait() // 4
	//	fmt.Println("main finished")

	var cc chan int = make(chan int, 1)
	cc <- 99
	fmt.Println(<-cc)

	cha := make(chan int, 10)
	go fibonacci(cap(cha), cha)
	v, ok := <-cha // ok == false is close and no data
	fmt.Println(v, ok)
	for i := range cha {
		fmt.Println(i)
	}

	c1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c1)
		}
		quit <- 0
	}()
	fibonacci_new(c1, quit)

	c2 := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c2:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
				//			default:
				//				fmt.Println("STUCK, NO SIGNAL")
			}
		}
	}()
	<-o
}
