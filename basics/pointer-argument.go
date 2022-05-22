package main

import "fmt"

func add(xPtr *int) {
	*xPtr = *xPtr + 10
}

func main() {
	var a int = 10
	//var aPtr *int = &a
	//add(aPtr) // pass by Pointer
	add(&a)
	fmt.Println(a)

	// 和使用者要求輸入，運用到指標參數 Pass by Pointer
	var msg string
	var msgPtr *string = &msg
	fmt.Scanln(msgPtr) // 傳遞字串變數的指標(記憶體位址)
	fmt.Println(msg)
}

/*
func add(x int) {
	x = x + 10
	fmt.Println("Add Function", x)
}

func main() {
	var a int = 10
	add(a) // pass by value
	fmt.Println("Main Function", a)
}
*/
