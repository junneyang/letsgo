package test

import (
	"fmt"
	"testing"
)

func max(a int64, b int64) (max int64) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func sumAndmul(a int64, b int64) (sum int64, mul int64) {
	sum = a + b
	mul = a * b
	//	return sum, mul
	return
}

func argsSum(args ...int64) (sum int64) {
	for _, v := range args {
		fmt.Println(v)
		sum += v
	}
	return
}

func getSlice() (result []int64) {
	result = append(result, 1)
	result = append(result, 2)
	result = append(result, 3)
	return result
}

func Test_Func(t *testing.T) {
	var a, b int64 = 10, 11
	t.Log("max: ", max(a, b))

	sum, mul := sumAndmul(a, b)
	t.Log("sum: ", sum)
	t.Log("mul: ", mul)

	t.Log("sumArgs: ", argsSum(1, 2, 3, 4, 5))

	// []  slice      [N/...]  array
	slice1 := []int64{1, 2, 3}
	slice2 := slice1
	s := append(slice2, 111, 222)
	s[0] = 888
	t.Log(slice1, len(slice1), cap(slice1))
	t.Log(slice2, len(slice2), cap(slice2))
	t.Log(s, len(s), cap(s))

	for i := 0; i < 5; i++ {
		defer t.Log(i)
	}

	slice := getSlice()
	t.Log(slice)
}
