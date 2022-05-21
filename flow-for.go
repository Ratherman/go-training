package main

import "fmt"

func main() {

	//基本迴圈使用
	/*var x int = 0
	for x < 3 {
		fmt.Println(x)
		x++
	}
	var x int
	for x = 0; x < 5; x += 2 {
		fmt.Println(x)
	}
	*/

	// 實際範例, 計算 1 + 2 + 3 + ... + 50 的結果

	// 彭彭的寫法
	var result int = 0
	var x int = 1
	for x <= 50 {
		result = result + x
		x++
	}
	fmt.Println(result)

	// 我的寫法
	/*
			var x int
			var result int = 0
			for x = 1; x <= 50; x++ {
				result += x
			}

		fmt.Println(result)
	*/
}
