package main

import (
	"encoding/json"
	"fmt"
	"log"
	"unicode/utf8"
)

type MyData struct {
	One int
	two string
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("DEFER: ", err)
		}
	}()
	//	x := 1
	//	fmt.Println(x) //prints 1
	//	{
	//		//		fmt.Println(x) //prints 1
	//		//		x := 2
	//		//		fmt.Println(x) //prints 2

	//		x = 2
	//		fmt.Println(x) //prints 2
	//	}
	//	fmt.Println(x) //prints 1 (bad if you need 2)

	var a []int64
	fmt.Println(a == nil)
	var b int64
	fmt.Println(b == 0)

	var c []int64
	fmt.Println(c)
	c = make([]int64, 0)
	c = append(c, 123)
	c = append(c, 456)
	c = append(c, 789)
	fmt.Println(c)
	fmt.Println(cap(c))
	fmt.Println(len(c))

	var d map[string]int64
	fmt.Println(d)
	d = make(map[string]int64)
	d["123"] = 123456
	fmt.Println(d)
	//	fmt.Println(cap(d))
	fmt.Println(len(d))

	var str string
	fmt.Println(str)
	fmt.Println(str == "")
	//	fmt.Println(str == nil)

	var slice []int64 = nil
	fmt.Println(slice)
	//	fmt.Println(slice == []int64)
	fmt.Println(slice == nil)

	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x)
	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(*arr) //prints [7 2 3]
	}(&x)
	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	xx := x[:]
	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(xx)
	fmt.Println(xx) //prints [1 2 3] (not ok if you need [7 2 3])

	y := []int64{12, 34, 78}
	for k, v := range y {
		fmt.Println(k, v)
	}

	array2 := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{6, 7, 8, 9}}
	fmt.Println(array2)

	ax := 2
	ay := 4
	table := make([][]int, ax)
	for i := range table {
		table[i] = make([]int, ay)
	}
	table[0][0] = 99
	table[1][3] = 88
	fmt.Println(table)

	mapa := map[string]int64{"aa": 123, "bb": 456}
	fmt.Println(mapa["aa"])
	fmt.Println(mapa["cc"])
	kk, vv := mapa["dd"]
	fmt.Println(kk, vv)
	if _, ok := mapa["ee"]; !ok {
		fmt.Println("NO ENTRY !")
	}

	strx := "卧槽ABCDEF"
	//	strx[0] = 'X'
	strx_slice := []rune(strx)
	//	strx_slice := []byte(strx)
	strx_slice[0] = '傻'
	stry := string(strx_slice)
	fmt.Println(stry)
	fmt.Println(strx)
	cc := strx[0]
	fmt.Println(string(cc))
	dd := strx_slice[0]
	fmt.Println(string(dd))

	fmt.Println(len("Hello, 世界"))
	fmt.Println(len([]rune("Hello, 世界")))

	temp := []byte("您好junneyang")
	fmt.Println(len(temp))
	fmt.Println(utf8.RuneCount(temp))
	fmt.Println(utf8.RuneCountInString("您好junneyang"))

	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true
	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false
	data3 := "您好junneyang"
	fmt.Println(utf8.ValidString(data3)) //prints: true
	for _, v := range data3 {
		fmt.Println("PPP: ", string(v))
	}
	for _, v := range []byte(data3) {
		fmt.Println("PPP: ", string(v))
	}
	for _, v := range []rune(data3) {
		fmt.Println("PPP: ", string(v))
	}

	xxx := []int{
		1,
		2, //error
	}
	fmt.Println(xxx)

	yyy := []int{3, 4} //no error
	fmt.Println(yyy)

	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
	mm := 1
	mm += 10
	fmt.Println(mm)

	fmt.Println(0x1 & 0x0)
	fmt.Println(0x1 | 0x0)
	fmt.Println(^0x1 + 1)

	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) //prints main.MyData{One:1, two:"two"}
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) //prints {"One":1}
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) //prints main.MyData{One:1, two:""}

	//	log.Fatalln("FATAL ERROR")
	log.Panicln("PANIC !!!")

}
