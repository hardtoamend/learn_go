package main

import "fmt"

// 常量
//const pi = 3.1415
//const e  = 2.7

// 批量常量
//const (
//	pi = 3.1415
//	e  = 2.7
//)

//const (
//	n1 = 10
//	n2 = 20
//	n3 = 30
//)

// iota 遇 count 为0 , 新增一行计数一次 , 相当于 const 索引
//const (
//	n1 = iota
//	n2
//	n3
//	n4
//)

// iota _ 跳过
//const (
//	n1 = iota
//	n2 = iota
//	_
//	n4 = iota
//)

// iota 中间插队
//const (
//	n1 = iota
//	n2
//	n3 = 100
//	n4 = iota
//)
//const n5 = iota

//func main() {
//	fmt.Println(n1, n2, n3, n4, n5)
//}

const (
	_  = iota
	KB = 1 << (10 * iota) // 1<<100
	MB
	GB
	TB
	PB
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
