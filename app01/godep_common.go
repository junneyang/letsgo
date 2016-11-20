package main

import (
	"common_c"
	"fmt"
)

func main() {
	c, err := common.Functest(12, 13)
	fmt.Println(c, err)
}
