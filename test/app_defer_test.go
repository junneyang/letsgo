package test

import (
	"testing"
)

func ftest(a int64, b int64) (result int64) {
	if a < 0 || b < 0 {
		panic("params invalid")
	}
	result = a * b
	return
}

func Test_Defer(t *testing.T) {
	t.Log(1)
	defer func() { t.Log("here3") }()
	defer func() {
		t.Log(4)
		if err := recover(); err != nil {
			t.Log(err)
		}
		t.Log(5)
	}()
	t.Log(2)

	//	panic("some exception")
	result := ftest(-10, 20)
	t.Log(result)

	defer func() { t.Log("here1") }()
	t.Log(3)
	defer func() { t.Log("here2") }()
}
