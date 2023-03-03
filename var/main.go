package main

import (
	"fmt"
)

// 全局变量
//var x = 100
//var y = "下"

func main() {
	// 标准声明
	var name string
	var age int
	var isok bool
	fmt.Println(name, age, isok)

	// 批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	// 声明变量同事指定初始值
	var name1 string = "小"
	var age1 int = 1
	fmt.Println(name1, age1)

	var name2, age2 = "哈", 2
	fmt.Println(name2, age2)

	// 类型推导 , 让编译器根据变量的初始值推导出其类型
	var n3 = "小"
	var a3 = "3"
	fmt.Println(n3, a3)

	// 短变量声明
	m := 10
	fmt.Println(m)

	// 匿名变量

}
