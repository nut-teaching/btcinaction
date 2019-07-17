package main

import "fmt"

type user struct {
	name string
	age  int
}

type student struct {
	user // 匿名嵌入其他结构类型
	school string
}

func main() {
	var s student

	fmt.Println(s)

	s.name = "nut"
	s.age = 18

	s.school = "清华"

	fmt.Println(s)
}
