package main

import (
	"fmt"
	//	"time"
)

func main() {
	ch := make(chan string)
	//	var ch chan string
	go func() {
		//		for m := range ch {
		//			fmt.Println("processed:", m)
		//		}
		//		fmt.Println(<-ch)
		//		fmt.Println(<-ch)

		ch <- "cmd.1"
		ch <- "cmd.2" //won't be processed
	}()
	//	ch <- "cmd.1"
	//	ch <- "cmd.2" //won't be processed
	//	time.Sleep(5 * time.Second)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
