package main

import (
	"fmt"
	"sort"
)

// 切片 slice
func main() {
	//// 1. 定义切片 slice
	//var a []string
	//var b []int
	//var c = []bool{false, true}
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//// 2. 基于数组定义切片 slice
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3] // s := a[low:high]
	//fmt.Printf("s = %v , s type = %T ,len(s): %v cap(s): %v \n", s, s, len(s), cap(s))

	////3. 切片再次切片
	//c := s[0:len(s)]
	//fmt.Printf("c = %v , c type = %T ,len(c): %v cap(c): %v \n", c, c, len(c), cap(c))
	//
	//// make 函数构造切片  第一个类型，第二个长度，第三个容量(上限), 第三个可以省略， 省略默认和长度一致
	//d := make([]int, 5, 10)
	//fmt.Printf("d = %v , d type = %T ,len(d): %v cap(d): %v \n", d, d, len(d), cap(d))
	//
	//// 通过 len 获取切片长度
	//fmt.Println("切片长度： ", len(d))
	//// 通过 cap 获取切片容量  最多可以存放多少个元素
	//fmt.Println("切片容量： ", cap(d))

	//// 4. nil
	//var a []int     // 声明 int 类型切片
	//var b = []int{} // 声明并且初始化
	//c := make([]int, 0)
	//if a == nil {
	//	fmt.Println("a == nil") // a == nil
	//}
	//if b == nil {
	//	fmt.Println("b == nil") // b != nil
	//}
	//if c == nil {
	//	fmt.Println("c == nil") // c != nil
	//}
	//fmt.Println(a, len(a), cap(a))
	//fmt.Println(b, len(b), cap(b))
	//fmt.Println(c, len(c), cap(c))

	//// 5. 切片赋值拷贝
	//a := make([]int, 3) // int[]{0,0,0}
	//b := a
	//b[0] = 100
	//fmt.Println(a)
	//fmt.Println(b)

	//// 6. 切片遍历
	//// 根据索引遍历
	//a := []int{1, 2, 3, 4, 5}
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("fori a = %d, a[%d] = %d \n", a, i, a[i])
	//}
	//fmt.Println()
	//// for range 遍历
	//for i, v := range a {
	//	fmt.Printf("forr a = %d, a[%d] = %d \n", a, i, v)
	//}

	// 7. 切片的扩容
	// 切片要初始化之后才能使用 使用 make 初始化 或者 定义的时候初始化
	//var a []int // 此时并没有申请内存没办法使用
	//a[0] = 100
	//fmt.Println(a)

	//var a []int // 此时并没有申请内存没办法使用
	//a = append(a, 10) // append 可能造成底层数组变化, 原来的存储不够存储到添加的时候, 会自动扩容 , 必须要用一个(原来切片)变量去接收返回
	//fmt.Println(a)

	//// for 循环 append 扩容
	//for i := 0; i < 10; i++ {
	//	a = append(a, i)
	//	fmt.Printf("%v , len(a) %d, cap(a) %d, prt: %p \n", a, len(a), cap(a), a)
	//}

	//// append 追加切片元素 将 b 追加到 a
	//a = append(a, 1, 2, 3, 4, 5)
	//fmt.Println(a)
	//b := []int{11, 12, 13, 14, 15}
	//fmt.Println(b)
	//a = append(a, b...)
	//fmt.Println(a)

	//// 8. copy() 函数 copy 切片
	//a := []int{1, 2, 3, 4, 5}
	//b := make([]int, 5, 5)
	//c := b
	//copy(b, a)
	//b[0] = 100
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//// 9. 切片删除元素
	//a := []string{"北京", "上海", "广州", "深圳"}
	//a = append(a[:2], a[3:]...)
	//fmt.Println(a)
	//
	//// append(a[:index], a[index+1:])	// 删除 index 元素
	//a = append(a[:1], a[1+1:]...)
	//fmt.Println(a)

	////1.请写出下面代码的输出结果.
	//var a = make([]string, 5, 10) // [     ]
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v ", i))
	//}
	//fmt.Println(a) // [     0  1  2  3  4  5  6  7  8  9 ]

	//2.请使用内置的 `sort` 包对数组 `var a = [...]int{3, 7, 8, 9, 1}` 进行排序 (附加题,自行查资料解答) .
	var a = [...]int{3, 7, 8, 9, 1}
	// a[:] 得到的是一个切片, 指向底层数组 a
	sort.Ints(a[:])
	fmt.Println(a)

	// 1. 定义切片 slice
	// 2. 基于数组定义切片 slice
	// 3. 切片再次切片
	// 4. nil
	// 5. 切片赋值拷贝
	// 6. 切片遍历
	// 7. 切片的扩容
	// 8. copy() 函数 copy 切片
	// 9. 切片删除元素

	////1.请写出下面代码的输出结果.
	//var a = make([]string, 5, 10) // [     ]
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v ", i))
	//}
	//fmt.Println(a) // [     0  1  2  3  4  5  6  7  8  9 ]
	//2.请使用内置的 `sort` 包对数组 `var a = [...]int{3, 7, 8, 9, 1}` 进行排序 (附加题,自行查资料解答) .
}
