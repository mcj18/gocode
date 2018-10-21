package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 10)
	fmt.Println(mySlice)
	fmt.Printf("%T\n", mySlice)
	fmt.Printf("%v\n", mySlice)

	for i := 0; i < 30; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("index: ", i, "*** ", "value: ", mySlice, "*** ", "len: ", len(mySlice))
	}
}
