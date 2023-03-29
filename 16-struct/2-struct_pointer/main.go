package main

import "fmt"

// 结构体指针
type person struct {
	name, city string
	age        int8
}

func main() {
	var p3 = new(person)
	fmt.Printf("p3 = %#v \n", p3)
	//(*p3).name = "小白"
	//(*p3).city = "北京"
	//(*p3).age = 23
	p3.name = "小白"
	p3.city = "北京"
	p3.age = 23
	fmt.Printf("%#v \n", p3)

	// 取结构体的地址进行实例化
	p4 := &person{}
	fmt.Printf("%T \n", p4)
	fmt.Printf("%#v \n", p4)
	p4.name = "小白"
	p4.city = "北京"
	p4.age = 18
	fmt.Printf("%#v \n", p4)
}
