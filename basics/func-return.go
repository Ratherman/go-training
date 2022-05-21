package main

import "fmt"

func show(msg string) {
	if msg == "Hello" {
		return
	}
	fmt.Println(msg)
}

func add(n1 int, n2 int) int {
	var result int = n1 + n2
	return result
}

func test() (int, string) {
	return 5, "Test"
}

func main() {

	// 多個回傳值
	var x int
	var s string
	x, s = test()
	fmt.Println(x, s)

	// 基本的 return 運作方式
	/*
		show("Hello")
		show("你好")
		show("測試")
	*/

	// 單一回傳值
	/*
		var x int = add(3, 4)
		fmt.Println(x * 20)
	*/
}
