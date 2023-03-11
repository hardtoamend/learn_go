package main

import (
	"fmt"
	"strings"
)

// map (映射)
func main() {

	//// 只声明 map 类型 ， 但是没有初始化 , a 就是初始值 nil
	//var a map[string]int
	//fmt.Println(a == nil)
	//
	//// map 初始化
	//a = make(map[string]int, 0)
	//fmt.Println(a == nil)
	//
	//// map 添加键值对
	//a["哇哦"] = 100
	//a["小红"] = 200
	//fmt.Println(a)
	//fmt.Printf("type: %T \n", a)
	//fmt.Printf("a: %#v \n", a)
	//
	//// 声明 map 的同时完成初始化
	//b := map[int]bool{
	//	1: true,
	//	2: false,
	//}
	//fmt.Printf("b: %#v \n ", b)
	//fmt.Printf("type: %T \n", b)

	////错误声明赋值 (为初始化 map ) , map 没有初始化不能直接操作 , 没有初始化在内存里没有位置 , 需要初始化申请内存地址才可以使用
	//var c map[int]int
	//c[100] = 200
	//fmt.Println(c)

	//// 判断某个键存不存在
	//var scoreMap = make(map[string]int, 8)
	//scoreMap["小红"] = 100
	//scoreMap["小蓝"] = 200
	//
	//// 判断 张二狗子 在不在 soreMap 中
	//v, ok := scoreMap["张二狗子"]
	//fmt.Printf("v: %v , ok: %v \n", v, ok)
	//if ok {
	//	fmt.Println("张二狗子在 scoreMap 中", v)
	//} else {
	//	fmt.Println("查无此人 ! ")
	//}
	//
	//value, ok := scoreMap["小红"]
	//fmt.Printf("value: %v , ok: %v \n", value, ok)
	//if ok {
	//	fmt.Printf("%v , %v \n", value, ok)
	//} else {
	//	fmt.Println("不存在")
	//}
	//
	////map 遍历 for range
	////键值对顺序跟添加顺序无关 , map 是无序的
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}
	//
	////无需 value 只需 key （只遍历 key）
	//for k, _ := range scoreMap {
	//	fmt.Println(k)
	//}
	//
	//// 无需 key 只需 map	(只遍历 value)
	//for _, v := range scoreMap {
	//	fmt.Println(v)
	//}
	//
	//// 删除键值对 小红
	//delete(scoreMap, "小红")
	//fmt.Println(scoreMap)
	//
	//// 按照固定顺序遍历 map
	//var numMap = make(map[string]int, 100)
	//
	//// 添加 50 个键值对
	//for i := 0; i <= 50; i++ {
	//	key := fmt.Sprintf("stu%02d", i)
	//	value := rand.Intn(100) // rand.Intn 0～99 生成随机整数
	//	numMap[key] = value
	//}
	//
	//for k, v := range numMap {
	//	fmt.Println(k, v)
	//}
	//
	//// 按照 key 从小到大的顺序遍历 map
	//// 1. 先取出所有的 key 存放到切片中
	////初始化定义切片
	//keys := make([]string, 0, 100)
	//
	//for k := range numMap {
	//	keys = append(keys, k)
	//}
	//
	////fmt.Println(keys[1])
	//// 2. 对 key 做排序
	//sort.Strings(keys) // keys 目前是一个有序的切片
	//// 3. 按照排序后的 key 对 numMap 排序
	//for _, key := range keys {
	//	//fmt.Println(key, numMap[key])
	//	fmt.Println(key, numMap[key])
	//}

	//// 复杂类型 map , 元素类型为 map 的切片
	//var mapSlice = make([]map[string]int, 8, 8) // 只是完成了 slice 的初始化
	//// [nil nil nil nil nil nil nil nil ]
	//fmt.Println(mapSlice[0] == nil)
	//// 还需要完成内部的 map 元素初始化
	//mapSlice[0] = make(map[string]int, 8) // 完成了 map 的初始化
	//mapSlice[0]["小红"] = 100
	//fmt.Println(mapSlice)

	//// 值 为 切片 的 map
	//var sliceMap = make(map[string][]int, 8)
	//v, ok := sliceMap["中国"]
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	// sliceMap 没有中国这个键
	//	sliceMap["中国"] = make([]int, 8, 8) // 完成切片初始化
	//	sliceMap["中国"][0] = 100
	//	sliceMap["中国"][1] = 200
	//	sliceMap["中国"][2] = 300
	//}
	//
	//// 遍历 sliceMap
	//for k, v := range sliceMap {
	//	fmt.Println(k, v)
	//}

	// 练习 统计一个字符串中每个单词出现的次数
	// "how do you do" 中每个单词出现的次数

	// 0. 定义一个 map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	// 1. 字符串中都有哪些单词
	words := strings.Split(s, " ")

	// 2. 遍历单词做统计
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1 // map 中有这个单词的统计记录
		} else {
			wordCount[word] = 1 // map 中没有这个单词的统计记录
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}

// 小结:
// 1. map 定义和初始化 (定义 map 的时候做初始化 , 也可以使用 make ) , map 如果没有初始化就是 nil , 无法直接使用 , 必须要经过初始化才能使用
// 2. map 遍历
// 3. 判断键在 map  中存不存在
// 4. 添加键值对
// 5. delete 删除键值对
// 6. 按照学生学号遍历 map
// 7. 元素为 mpa 类型的切片 (无论 sclice 还是map , 引用类型都要初始化之后才可以使用, 使用 make 初始化的时候一定要提前规划好 容量, 减少后续代码运行过程中动态扩容或者申请内存)
// 8. map 值为 切片 类型   (无论 sclice 还是map , 引用类型都要初始化之后才可以使用, 使用 make 初始化的时候一定要提前规划好 容量, 减少后续代码运行过程中动态扩容或者申请内存)
