package main

import (
	"fmt"

	"github.com/junneyang/letsgo/common"
)

func main() {
	c, err := common.Functest(12, 13)
	fmt.Println(c, err)
}
