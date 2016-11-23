package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/junneyang/letsgo/common"
	"github.com/pkg/errors"
)

func FuncTestWrap() (c int64, err error) {
	a, err := strconv.ParseInt(os.Getenv("VAR_A"), 10, 64)
	if err != nil {
		return c, errors.Wrap(err, "VAR_A ParseInt Fail")
	}

	b, err := strconv.ParseInt(os.Getenv("VAR_B"), 10, 64)
	if err != nil {
		return c, errors.Wrap(err, "VAR_B ParseInt Fail")
	}

	c, err = common.Functest(a, b)
	return c, errors.Wrap(err, "FuncTestWrap Fail")
}

func main() {
	c, err := FuncTestWrap()
	fmt.Println(c, err)
}
