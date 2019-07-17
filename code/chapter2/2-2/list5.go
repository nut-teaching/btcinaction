package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 9; i >= 0; i-- {
		fmt.Println(i)
	}

	// 相当于 while (x < 10)
	var x int
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// 相当于 while (true)
	var y int
	for {
		fmt.Println(y)
		y++

		if y == 10 {
			break
		}
	}

	arr := []int{1, 2, 3, 4, 5,}

	for i, a := range arr {
		fmt.Println(i, a)
	}
}
