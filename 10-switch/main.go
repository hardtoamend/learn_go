package main

func main() {
	//// 1. case 基本使用
	//finger := 3
	//switch finger {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小拇指")
	//default:
	//	fmt.Println("无效输入")
	//}

	//// 2. case 一次判断多个
	//num := 5
	//switch num {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("奇数")
	//case 2, 4, 6, 8, 10:
	//	fmt.Println("偶数")
	//}

	//// 3. case 中使用biaodashi
	//age := 30
	//switch {
	//case age > 18:
	//	fmt.Println("澳门首家线上赌场开业了")
	//case age < 18:
	//	fmt.Println("waring")
	//default:
	//	fmt.Println("欧吼")
	//}

	//// 4. goto 跳转到指定标签 for 循环嵌套
	//var breakFlag bool
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			// 设置退出标签
	//			breakFlag = true
	//			break
	//		}
	//		fmt.Printf("%v-%v\n", i, j)
	//	}
	//	// 外层for循环判断
	//	if breakFlag {
	//		break
	//	}
	//}

	//	// 4. goto 跳转到指定标签
	//	for i := 0; i < 10; i++ {
	//		for j := 0; j < 10; j++ {
	//			if j == 2 {
	//				// 设置退出标签
	//				goto breakTag
	//			}
	//			fmt.Printf("%v-%v\n", i, j)
	//		}
	//	}
	//	return
	//	// 标签
	//breakTag:
	//	fmt.Println("结束for循环")
}
