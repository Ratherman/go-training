package main // 可執行程式必須使用 main 封包
import "fmt" // 載入內建的 fmt 封包，用來做基本的輸出輸入
func main(){ // 建立 main 函式，程式的進入點
	// 執行程式時，從這邊開始
	// 輸出訊息到終端: fmt.Println(資料,資料,...)
	/*
	fmt.Println(3) // 整數 int
	fmt.Println(3.1415) // 浮點數 float64
	fmt.Println("測試中文") // 字串 string
	fmt.Println(true) // 布林值 bool
	fmt.Println('a') // 字符 rune
	*/
	// 變數的使用
	var x int // 宣告變數 x
	x = 4 // 指定資料：把資料 4 放進變數裡
	fmt.Println(x) // 使用變數：用變數的名稱代替資料，做操作
	x=10 // 指定新的資料：把資料 10 放進變數，覆蓋原來的資料
	fmt.Println(x)
	// x="Hello" 資料型態不正確
	var f float64=3.1415 // 宣告變數訓便放進資料： var 變數名稱 資料型態=適當的資料
	fmt.Println(f)
	var s string="哈囉 Golang"
	fmt.Println(s)
	var test bool=true
	fmt.Println(test)
	var c rune='b'
	fmt.Println(c)
}