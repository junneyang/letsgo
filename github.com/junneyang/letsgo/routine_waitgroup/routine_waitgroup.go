package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, wg)
	}
	wg.Wait()
	fmt.Println("all done!")
}
func doit(workerId int, wg sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n", workerId)
}
