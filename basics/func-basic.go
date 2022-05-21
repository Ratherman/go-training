package main

import "fmt"

func test(msg string) {
	fmt.Println(msg)
}

func add(n1 int, n2 int) {
	var result int = n1 + n2
	fmt.Println(result)
}

func sum(max int) {
	var result = 0
	var n int
	for n = 1; n <= max; n++ {
		result += n // result = result + n
	}
	fmt.Println(result)
}

func main() {

	// 計算 1+2+3+...+max 的結果
	sum(100) // 1+2+3+...+100
	sum(50)  // 1+2+3+...+50
	sum(10)  // 1+2+3+...+10
	// 基礎函式語法練習
	/*
		test("Hello")
		test("你好")
		add(3, 4)
		add(-2, 5)
	*/

}
