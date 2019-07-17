package main

import "fmt"

/*
	用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用
	常用来释放资源、解除锁定，或执行一些清理操作
*/

func testFunc() {
	defer fmt.Println("release")
	defer fmt.Println("release1")
	defer fmt.Println("release2")

	fmt.Println("this is test function")
}
func main() {

	testFunc()
}
