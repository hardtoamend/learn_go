package main

import "fmt"

// 数组相关内容
func main() {
	//// 未声明数组值 , 默认为 0
	//var a [3]int
	//var b [4]int
	//fmt.Println(a) // [0 0 0]
	//fmt.Println(b) // [0 0 0 0]
	//
	//// 数组初始化
	//// 1.  定义时使用初始值列表的方式初始化
	//var cityArray = [4]string{"上海", "北京", "广州", "深圳"}
	//fmt.Println(cityArray)
	//fmt.Println(cityArray[3])
	//fmt.Println(cityArray[1])
	//
	//// 2. 编译器推导数组长度
	//var boolArray = []bool{true, false, true, false}
	//fmt.Println(boolArray)
	//
	//// 3. 使用索引值方式初始化
	//var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	//fmt.Println(langArray)
	//fmt.Printf("langArray Type : %T \n", langArray)

	// 数组的遍历
	var cityArray = [4]string{"上海", "北京", "广州", "深圳"}
	for i := 0; i < len(cityArray); i++ {
		fmt.Printf("[%d]string %s\n", i, cityArray[i])
	}

	// for range 遍历
	// 获取 Array index 索引 & value 值
	for index, value := range cityArray {
		fmt.Println(index, value)
	}
	// 获取 Array index 索引
	for index := range cityArray {
		fmt.Println(index)
	}
	// 获取 Array value 值
	for _, value := range cityArray {
		fmt.Println(value)
	}
}
