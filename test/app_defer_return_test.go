package test

import (
	"testing"

	"github.com/pkg/errors"
)

func Functest(a int64, b int64) (c int64, err error) {
	if a < 0 || b < 0 {
		return c, errors.New("a or b is invalid")
	}
	c = a * b
	return c, err
}

func Test_Defer_Return(t *testing.T) {
	defer func() {
		println("XXXX")
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	_, err := Functest(-1, 2)
	if err != nil {
		println(err.Error())
		return
	}
}
