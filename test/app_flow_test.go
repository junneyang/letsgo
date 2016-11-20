package test

import (
	//	"fmt"
	"testing"
)

func Test_FlowControl(t *testing.T) {
	var x int64 = 10
	if x == 10 {
		// fmt.Println("x is  10")
		t.Log("x is  10")
	} else if x >= 10 {
		t.Log("x is  >=10")
	} else {
		// fmt.Println("x is not 10")
		t.Log("x is not 10")
	}
	t.Log(x)

	if temp := 10; temp == 10 {
		t.Log("temp == 10")
		t.Log(temp)
	}
	//	t.Log(temp)

	i := 0
LOOP:
	//	fmt.Println(i)
	t.Log(i)
	i++
	if i <= 100 {
		goto LOOP
	} else {
		return
	}
}
