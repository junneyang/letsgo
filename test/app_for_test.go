package test

import (
	//	"fmt"
	"testing"
)

func Test_For(t *testing.T) {
	//	sum := 0
	//	for i := 0; i <= 5; i++ {
	//		t.Log("i: ", i)
	//		sum += i
	//	}
	//	t.Log("sum: ", sum)

	sum := 0
	for i, j := 0, 0; i <= 5 && j <= 5; i, j = i+1, j-1 {
		t.Log("i: ", i)
		t.Log("j: ", j)
		sum += i + j
	}
	t.Log("sum: ", sum)

	//	var i, j int64 = 0, 0

	//	for true {
	//		t.Log("HAHA")
	//	}

	count := 0
	for count <= 100 {
		count += 1
	}
	t.Log("count: ", count)

	for index := 10; index > 0; index-- {
		if index == 5 {
			//			break
			continue
		}
		//		fmt.Println(index)
		t.Log(index)
	}

	map1 := make(map[string]string)
	map1["a"] = "AAA"
	map1["b"] = "BBB"
	map1["c"] = "CCC"
	for k, v := range map1 {
		t.Log(k, v)
	}
	for _, v := range map1 {
		t.Log(v)
	}

	array := [...]int64{1, 2, 3, 4}
	for k, v := range array {
		t.Log(k, v)
	}
	for _, v := range array {
		t.Log(v)
	}

	slice := array[:3:3]
	for k, v := range slice {
		t.Log(k, v)
	}
	t.Log(cap(slice))
	slice[0] = 444
	slice = append(slice, 888)
	slice = append(slice, 999)
	t.Log(slice)
	t.Log(array)

}
