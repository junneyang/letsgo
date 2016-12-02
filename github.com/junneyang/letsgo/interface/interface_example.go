package main

import (
	"fmt"
	//	"path"
	//	"runtime"
	//	"strings"

	"github.com/pkg/errors"
)

type Rectangle struct {
	width  float64
	height float64
}

type ERROR struct {
	pkg  string
	info string
	err  error
}

func (e *ERROR) Error() string {
	if (*e).err == nil {
		return fmt.Sprintf("%s: %s", (*e).pkg, (*e).info)
	} else {
		return fmt.Sprintf("%s: %s, %v", (*e).pkg, (*e).info, (*e).err)
	}
}

func (r Rectangle) area() (area float64) {
	return r.width * r.height
}

//func (r *Rectangle) area() (area float64) {
//	return (*r).width * (*r).height
//}

func (r *Rectangle) setAttr(width float64, height float64) {
	r.height = height
	r.width = width
}

func (r Rectangle) setAttr2(width float64, height float64) {
	r.height = height
	r.width = width
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("DEFER: ", p)
			//			if e, ok := p.(error); ok {
			//				fmt.Println("DEFER: ", e.(pkg))
			//				fmt.Println("DEFER: ", e.(info))
			//				fmt.Println("DEFER: ", e.(err))
			//			}
			e, ok := p.(error)
			fmt.Println(e, ok)
			fmt.Printf("%#v\n", p.(error))
			fmt.Printf("%#v\n", p.(error).Error())
		}
	}()
	r := Rectangle{11.1, 22.2}
	fmt.Println(r.area(), 11.1*22.2)

	rr := new(Rectangle)
	(*rr).width = 33.3
	(*rr).height = 44.4
	fmt.Println(rr.area(), 33.3*44.4)
	rr.width = 55.5
	rr.height = 66.6
	fmt.Println(rr.area(), 55.5*66.6)

	rr.setAttr(12, 21)
	fmt.Println(rr)
	fmt.Println(rr.area(), 12*21)

	rr.setAttr2(33, 44)
	fmt.Println(rr)
	fmt.Println(rr.area(), 33*44)

	//	pc, file, line, _ := runtime.Caller(2)
	//	_, fileName := path.Split(file)
	//	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	//	pl := len(parts)
	//	packageName := ""
	//	funcName := parts[pl-1]
	//	if parts[pl-2][0] == '(' {
	//		funcName = parts[pl-2] + "." + funcName
	//		packageName = strings.Join(parts[0:pl-2], ".")
	//	} else {
	//		packageName = strings.Join(parts[0:pl-1], ".")
	//	}

	//	fmt.Println(fileName)
	//	fmt.Println(packageName)
	//	fmt.Println(funcName)
	//	fmt.Println(line)

	packageName := "testPackage"
	err := errors.New("someException")
	panic(&ERROR{packageName, "HUCK!", err})
}
