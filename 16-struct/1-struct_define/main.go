package main

import "fmt"

// 定义结构体

type person struct {
	name, city string
	age        int8
}

// 结构体实例化 ， 实例化以后才可以分配内存使用
func main() {
	var p1 person
	p1.name = "小白"
	p1.city = "北京"
	p1.age = 23
	fmt.Printf("p1 = %#v \n", p1)
	fmt.Println(p1.city)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	// 匿名结构体
	p2 := struct {
		name    string
		married bool
	}{}
	p2.name = "小白"
	p2.married = false
	fmt.Printf("p2 = %#v \n", p2)
}
