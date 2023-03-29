package main

import "fmt"

// 结构体初始化
type person struct {
	name, city string
	age        int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	// 1. 键值对初始化
	p5 := person{
		name: "小白",
		age:  23,
	}
	fmt.Printf("%T \n", p5)
	fmt.Printf("%#v \n", p5)

	p6 := &person{
		name: "小白",
		city: "北京",
		age:  23,
	}
	fmt.Printf("%#v \n", p6)

	// 2. 值的列表进行初始化
	p7 := person{
		"小白",
		"北京",
		23,
	}
	fmt.Printf("%#v \n", p7)

	// 内存地址
	cacheAdd := person{
		name: "小白",
		city: "北京",
		age:  23,
	}
	fmt.Printf("cacheAdd.name %p %v \n", &cacheAdd.name, cacheAdd.name)
	fmt.Printf("cacheAdd.city %p %v \n", &cacheAdd.city, cacheAdd.city)
	fmt.Printf("cacheAdd.age %p %v \n ", &cacheAdd.age, cacheAdd.age)

	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p \n", &n.a)
	fmt.Printf("n.b %p \n", &n.b)
	fmt.Printf("n.c %p \n", &n.c)
	fmt.Printf("n.d %p \n", &n.d)
}
