package main

import "fmt"

// 1. 函数
// 1.1.1 定义一个不需要参数也没有返回值的函数: sayhello
func sayhello() {
	fmt.Println("Hello , World !!! ")
}

// 1.1.2 定一个一个接受 string 类型的 name 参数
func sayhello2(name string) {
	fmt.Println("Hello", name)
}

// 1.1.3 定义接收多个参数 并且有一个返回值
func intSum(x, y int) int {
	ret := x + y
	return ret
}

func intSum2(x, y int) (ret int) {
	ret = x + y
	return
}

// 1.1.4 函数接收可变参数 , 在参数名后 + ... 表示可变参数
// 可变参数在函数体中是切片类型
func intSum3(a ...int) (ret int) {
	//fmt.Println(a)
	//fmt.Printf("a type: %T \n", a)
	ret = 0
	for _, v := range a {
		ret = ret + v
	}
	return
}

// 1.1.5 固定参数 和 可变参数同时出现 , 可变参数放在最后
// go 语言中函数没有默认参数
func intSum4(a int, b ...int) (ret int) {
	//fmt.Println(a)
	//fmt.Printf("a type: %T \n", a)
	ret = a
	for _, v := range b {
		ret = ret + v
	}
	return
}

// Go 语言中函数参数类型简写
func intSum5(a, b int) (ret int) {
	ret = a + b
	return
}

// 1.1.6 定义具有多个返回值的函数 , 多返回值也支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	// 函数调用
	//sayhello()
	//sayhello2("哈哈")
	//r := intSum(1, 2)
	//fmt.Println(r)
	//
	//r1 := intSum3()
	//r2 := intSum3(1, 2)
	//r3 := intSum3(1, 2, 3)
	//fmt.Println(r1)
	//fmt.Println(r2)
	//fmt.Println(r3)
	//
	//re1 := intSum4(0)
	//re2 := intSum4(10, 20)
	//re3 := intSum4(10, 20, 30)
	//fmt.Println(re1)
	//fmt.Println(re2)
	//fmt.Println(re3)

	x, y := calc(100, 200)
	fmt.Println(x, y)
}
