package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int32
	var s = "i am string"

	fmt.Println(x, s)

	t := reflect.TypeOf(s)
	tx := reflect.TypeOf(x)

	fmt.Println(t.String(), tx.String())
}
