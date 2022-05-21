package main

import "fmt"

func main() {
	// break
	/*
		var x int = 0
		for x < 5 {
			if x == 3 {
				break
			}
			fmt.Println(x)
			x++
		}
	*/

	// continue
	/*
		var x int
		for x = 0; x < 5; x++ {
			if x%2 == 0 { // x 是偶數
				continue
			}
			fmt.Println(x)
		}
	*/

	// 實際範例，持續讓使用者輸入數字，計算總和。直到使用者輸入0為止
	var result int = 0

	for true { // 無窮迴圈

		var user_input int

		fmt.Scanln(&user_input)
		if user_input == 0 {
			break
		}
		result += user_input
	}
	fmt.Println("總和", result)
}
