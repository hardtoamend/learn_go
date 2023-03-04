package main

import "fmt"

// Go 语言中的运算符
func main() {
	//// 1. 算数运算符
	//a := 10
	//b := 20
	//fmt.Println(a + b)
	//fmt.Println(a - b)
	//fmt.Println(a * b)
	//fmt.Println(5 / 2) // 2
	//fmt.Println(5 % 2) // 1
	//a++
	//fmt.Println(a)
	//a--
	//fmt.Println(a)

	//// 2. 关系运算符
	//fmt.Println(10 > 2)
	//fmt.Println(10 != 2)
	//fmt.Println(4 <= 2)

	//// 逻辑运算符
	//fmt.Println(10 > 5 && 1 == 1) // && 逻辑AND 两边为 True 则为 True      否则为 False
	//fmt.Println(10 > 5 || 1 > 1)  // || 逻辑OR  两边有一边为 True 则为 True 否则为 False
	//fmt.Println(!(5 > 10))        // ！  逻辑NOT 如果条件为 True 则为 False  否则为 True

	//// 位运算符
	//a := 1             // 001
	//b := 5             // 101
	//fmt.Println(a & b) // 001 => 1
	//fmt.Println(a | b) // 001 => 5
	//fmt.Println(a ^ b) // 100 => 4
	//
	//fmt.Println(1 << 2) // 100 => 4
	//fmt.Println(4 >> 2) // 001 => 1

	// 赋值运算符
	var a int
	a = 5
	a += 5 // a = a + 5
	fmt.Println(a)

}
