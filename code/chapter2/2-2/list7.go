package main

import "fmt"

/*
	函数是第一类型，可作为参数或返回值
*/
func print(s string) func() { // 返回函数类型

	return func() { // 匿名函数
		fmt.Println(s) // 闭包
	}
}

func main() {

	f := print("hello")

	f()
}
