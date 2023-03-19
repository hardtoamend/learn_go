package main

import "fmt"

func main() {
	var a *int // a = nil
	a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["Hello"] = 100
	fmt.Println(b)
}
