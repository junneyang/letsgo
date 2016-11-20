package common

import (
	"fmt"

	"github.com/pkg/errors"
)

func Functest(a int64, b int64) (c int64, err error) {
	if a < 0 || b < 0 {
		c = -1
		err := errors.New("whoops")
		fmt.Println(err)
	} else {
		c = a * b
		fmt.Println(c)
	}
	return
}
