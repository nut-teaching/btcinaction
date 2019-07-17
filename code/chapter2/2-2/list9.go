package main

import "fmt"

/*
	切片（slice）可实现类似动态数组的功能
*/

func main() {
	x := make([]int, 0, 4) // 创建一个长度和容量都是0的切片
	fmt.Println(len(x), cap(x))
	for i := 0; i < 10; i++ {
		x = append(x, i) // 动态追加数据，当超出容量限制时，自动分配更大的存储空间
	}

	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i)
	}
	fmt.Println(len(y), cap(y))
}
