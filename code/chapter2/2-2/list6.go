package main

import (
	"errors"
	"fmt"
)

/*
	函数可定义多个返回值，甚至对其命名

 */

func add(x, y int) int {

	return x + y
}

func add2(x, y int) (r int) {

	r = x + y

	return
}

func div(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("division by zero ")
	}

	return x / y, nil

}

func div1(x, y int) (r int, err error) {

	if y == 0 {
		r = 0
		err = errors.New("division by zero ")

		return
	}

	r = x / y

	return
}

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(add2(1, 2))

	fmt.Println(div(4, 2))
	fmt.Println(div(4, 0))

	fmt.Println(div1(4, 2))
	fmt.Println(div1(4, 0))
}
