//package main

//import (
//	"bytes"
//	"encoding/binary"
//	"fmt"
//)

//func IntToBytes(n int) []byte {
//	x := int32(n)

//	bytesBuffer := bytes.NewBuffer([]byte{})
//	binary.Write(bytesBuffer, binary.BigEndian, x)
//	return bytesBuffer.Bytes()
//}

//func main() {
//	b := append([]byte("HEADER"), IntToBytes(len("HELLO"))...)
//	fmt.Println(b)
//}
package main

import (
	"fmt"
	"time"
)

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	// 将计算结果发送到channel中
	resultChan <- sum
}

func test(ch chan int) {
	for _ = range time.Tick(1e10) {
		fmt.Printf("cur conn num: %v\n", time.Now())
	}
	ch <- 1
}

func main() {
	//	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//	resultChan := make(chan int, 3)
	//	go sum(values[:len(values)/2], resultChan)
	//	go sum(values[len(values)/2:], resultChan)
	//	go sum(values[len(values)/3:], resultChan)
	//	sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
	//	fmt.Println("Result:", sum1, sum2, sum3)
	ch := make(chan int)
	go test(ch)
	<-ch
}
