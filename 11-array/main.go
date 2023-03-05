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

	//// 4. 数组的遍历
	//var cityArray = [4]string{"上海", "北京", "广州", "深圳"}
	//for i := 0; i < len(cityArray); i++ {
	//	fmt.Printf("[%d]string %s\n", i, cityArray[i])
	//}
	//
	//// 5. for range 遍历
	//// 获取 Array index 索引 & value 值
	//for index, value := range cityArray {
	//	fmt.Println(index, value)
	//}
	//// 获取 Array index 索引
	//for index := range cityArray {
	//	fmt.Println(index)
	//}
	//// 获取 Array value 值
	//for _, value := range cityArray {
	//	fmt.Println(value)
	//}

	//// 6. 二维数组
	//cityArray := [4][2]string{
	//	{"北京", "西安"},
	//	{"上海", "杭州"},
	//	{"重庆", "成都"},
	//	{"广州", "深圳"},
	//}
	////fmt.Println(cityArray)
	////fmt.Println(cityArray[2][0])		// 重庆
	//
	//// 7. 二维数组遍历
	//for _, v1 := range cityArray {
	//	for _, v2 := range v1 {
	//		fmt.Println(v2)
	//	}
	//}

	//// 8. 数组是值类型
	//x := [3][2]int{
	//	{1, 2},
	//	{3, 4},
	//	{5, 6},
	//}
	//fmt.Println(x)
	//f1(x)
	//fmt.Println(x)

	//// 1.求数组[1, 3, 5, 7, 8]所有元素的和

	//// 1.1 方法1
	//var sumarray = []int{1, 3, 5, 7, 8}
	//var sum int
	//for i := 0; i < len(sumarray); i++ {
	//	//fmt.Println(sumarray[i])
	//	sum += sumarray[i]
	//}
	//fmt.Println(sum)

	//// 1.2 方法2
	//for _, value := range sumarray {
	//	sum = sum + value
	//}
	//fmt.Println(sum)

	// 2. 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	var sumarray = []int{1, 3, 5, 7, 8}
	//for i := 0; i < len(sumarray); i++ {
	//	for j := 0; j < len(sumarray); j++ {
	//		if sumarray[i]+sumarray[j] == 8 {
	//			fmt.Printf("和为 8 的两个元素的下标为 (%d,%d) \n", i, j)
	//		}
	//	}
	//}

	for indexv1, valuev1 := range sumarray {
		for indexv2, valuev2 := range sumarray {
			if valuev1+valuev2 == 8 {
				fmt.Printf("和为 8 的两个元素的下标为 (%d,%d) \n", indexv1, indexv2)
			}
		}
	}

}

////8. 数组是值类型
//func f1(a [3][2]int) {
//	a[0][1] = 100
//}
