package main

import "fmt"

// 字符 byte & rune
func main() {
	// byte uint8 别名 , ADCII 码
	// rune int32 别名

	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T \n", c1, c2)

	s := "hello哇哦"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c \n", s[i])
	}

	for _, r := range s {
		fmt.Printf("%c \n", r)
	}
}
