package main

import (
	"common"
	"fmt"
)

func main() {
	c, err := common.Functest(12, 13)
	fmt.Println(c, err)
}
