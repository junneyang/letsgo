package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var p Person = Person{"aaa", 123}
	fmt.Println(p)
	p.name = "junneyang"
	p.age = 28
	fmt.Println(p)
	p = Person{name: "test", age: 456}
	fmt.Println(p)
	var pp *Person = new(Person)
	(*pp).name = "aha"
	(*pp).age = 22
	pp.age += 40
	fmt.Println(*pp)
	//	fmt.Println(pp)

	//	var qq Person = make(Person)
	//	fmt.Println(qq)
}
