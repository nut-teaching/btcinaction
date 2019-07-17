package main

import "fmt"

func main() {
	x := 1

	switch {
	case x > 0:
		fmt.Println("greater than zero")
	case x < 0:
		fmt.Println("lower than zero")
	default:
		fmt.Println("is zero")
	}

	switch x {
	case 1:
		fmt.Println("i am 1")
	case 0:
		fmt.Println("i am 0")
	case -1, -2:
		fmt.Println("i am -1 or -2")
	default:
		fmt.Println("i do not know who am i")

	}
}
