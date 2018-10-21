package main

import (
	"fmt"
	"strconv"
)

func main () {
	x := "1"
	y := 6

	fmt.Println(x)
	fmt.Println(y)

	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)
	fmt.Println([]byte("hello"))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", int(val))
	fmt.Printf("%T\n", val.(int))
}