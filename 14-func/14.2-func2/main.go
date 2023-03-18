package main

import (
	"fmt"
)

// 函数进阶之 变量作用域

// 定义全局变量 num
var num = 10

// 定义函数
func testGlobal() {
	num := 100
	name := "Hello"
	// 函数中使用变量
	// 1. 先在自己函数中查找 , 找到就用函数中
	// 2. 函数中找不到就往外层找全局变量
	fmt.Println("Global var = ", num)
	fmt.Println(name)
}

// 定义 add 函数 int 类型的参数 x,y 返回 int类型
func add(x, y int) int {
	return x + y
}

// 定义 sub
func sub(x, y int) int {
	return x - y
}

// 定义 calc 函数 int 类型参数 x,y , 变量为 op 的 func 参数 内int 参数 2 个 , 返回 int
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	//testGlobal()

	// 外层不能访问到内层函数的内部变量 (局部变量)
	//fmt.Println(name)

	// 函数可以作为变量
	//abc := testGlobal
	//fmt.Printf("abc type : %T \n", abc)
	//abc()

	r1 := calc(100, 200, add)
	fmt.Println(r1)

	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}
