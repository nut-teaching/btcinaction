package main

import "fmt"

func main() {
	x := 1

	if x > 0 {
		fmt.Println("greater than zero")
	} else if x < 0 {
		fmt.Println("lower than zero")
	} else {
		fmt.Println("is zero")
	}
}
