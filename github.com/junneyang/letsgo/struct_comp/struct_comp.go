package main

import (
	"fmt"
	//	"strconv"
	"reflect"
)

type IntfTest interface {
	SayHello()
	RetSomeStr() (result string)
}

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

type Employee struct {
	Human
	salary float64
}

func (h Human) SayHello() {
	fmt.Println("SayHello: ", h.name, h.age, h.weight)
}

func (h Human) RetSomeStr() (result string) {
	return "RetSomeStr: " + h.name
}

func (e Employee) SayHello() {
	fmt.Println("SayHello: ", e.name, e.age, e.weight, e.salary)
}

func (e Employee) RetSomeStr() (result string) {
	return "RetSomeStr: " + e.name + "----" + fmt.Sprintf("%v", e.salary)
}

func main() {
	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)

	mark.Human = Human{"aaa", 20, 123}
	mark.Human.age += 10
	fmt.Println("mark: ", mark)

	var i IntfTest
	i = Human{"JunneYang", 25, 120}
	i.SayHello()
	fmt.Println(i.RetSomeStr())

	i = Student{Human{"Stu", 25, 120}, "ENG"}
	i.SayHello()
	fmt.Println(i.RetSomeStr())

	i = Employee{Human{"Emp", 25, 120}, 8888.8}
	i.SayHello()
	fmt.Println(i.RetSomeStr())

	v, ok := i.(Employee)
	fmt.Println(v.salary, ok)

	//	var ttt interface{}
	switch ttt := i.(type) {
	case Employee:
		fmt.Printf("%T\n", ttt)
	case Human:
		fmt.Printf("%T\n", ttt)
	case Student:
		fmt.Printf("%T\n", ttt)
	default:
		fmt.Println("Unknown Type")
	}

	t := reflect.ValueOf(Human{}).Type()
	//	h := reflect.New(t).Elem()
	// new return address pointer
	h := reflect.New(t).Interface()
	fmt.Println(h)
	hh := h.(*Human)
	fmt.Println(hh)
	hh.SayHello()
	hh.age = 123
	hh.name = "abc"
	hh.weight = 345
	hh.SayHello()

	i = Human{"Emp", 25, 120}
	fmt.Println(reflect.TypeOf(i).Field(0).Type)
	fmt.Println(reflect.ValueOf(i).Field(1))
	//	reflect.ValueOf(i).Field(1).Elem().SetInt(88)
	//	fmt.Println(reflect.ValueOf(i).Field(1))
}
