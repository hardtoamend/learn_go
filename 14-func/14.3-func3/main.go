package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包

// 定义一个函数它的返回值是一个函数
// 把函数作为返回值
func a() func() {
	name := "小蓝"
	return func() {
		fmt.Println("Hello", name)
	}
}

// 示例 2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 示例 3
func cale(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 闭包 = 函数 + 外层变量
func main() {
	//sayHello := func() {
	//	fmt.Println("匿名函数")
	//}

	//func() {
	//	fmt.Println("匿名函数")
	//}()

	//sayHello()
	r := a()
	r()

	r1 := makeSuffixFunc(".txt")
	ret := r1("小红")
	fmt.Println(ret)

	r2 := makeSuffixFunc(".avi")
	ret2 := r2("小蓝")
	fmt.Println(ret2)

	x, y := cale(100)
	ret3 := x(200)
	fmt.Println(ret3) // base = 100 + 200
	ret4 := y(200)
	fmt.Println(ret4) // base = 300 - 200
}
