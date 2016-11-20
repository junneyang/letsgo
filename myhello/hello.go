package hello

import (
	"fmt"
)

func SayHello(name string) {
	fmt.Println("Hello, My name is: ", name)
	fmt.Printf("Hello, My name is: %s\n", name)
	fmt.Printf("Hello, My name is: %v\n", name)
}
