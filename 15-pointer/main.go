package main

import "fmt"

// 指针 pointer
// & 取址 * 取值	T 类型 *T 指针类型

/*
func main() {
	a := 10 // 定义变量 int类型a

	b := &a                       // b 求 a 的地址 *int 类型
	fmt.Printf("%v %p \n", a, &a) //10 0xc0000ac008
	fmt.Printf("%v %p \n", b, b)  // 0xc0000ac008 0xc0000ac008
	// 变量 b 是一个 int 类型的指针 （*int） ， 它存在的是变量 a 的地址

	// 根据 b 内存地址 取值
	c := *b
	fmt.Println(c)	// 10
}
*/

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)

}
