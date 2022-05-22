package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person = Person{"強強", 18}
	fmt.Println(p1.name, p1.age)
	var p2 Person = Person{age: 30, name: "圈圈"}
	p2.name = "小明"
	fmt.Println(p2.name, p2.age)
}
