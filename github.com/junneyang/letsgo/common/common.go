package common

import (
	"github.com/pkg/errors"
)

func Functest(a int64, b int64) (c int64, err error) {
	if a < 0 || b < 0 {
		return c, errors.New("a or b is invalid")
	}
	c = a * b
	return c, err
}
