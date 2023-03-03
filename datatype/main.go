package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制打印二进制
	n := 10
	fmt.Printf("%b \n", n)
	fmt.Printf("%d \n", n)

	// 八进制
	m := 075
	fmt.Printf("%o \n", m)
	fmt.Printf("%d \n", m)

	// 十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x \n", f)

	// unit8 (0-255)
	var age uint8
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)

	// 布尔值
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// 字符串 (双引号)
	s1 := "hello"
	s2 := "哈哈"
	fmt.Println(s1, s2)
	// 打印 win 路径 c:\code\go.exe
	fmt.Println("c:\\code\\go.exe")
	fmt.Println("\t 制表符 \n 换行符")

	// `` 内原样输出
	s3 := `
	test01
	哈哈哈
	aaa
	\t
	\n
	\r
`
	fmt.Println(s3)

}
