package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	var mySlice2 []int
	mySlice3 := make([]int, 50, 100)

	mySlice2 = append(mySlice2, 1, 2, 3, 4)

	fmt.Printf("%T\n", mySlice)
	fmt.Printf("%v\n", mySlice)
	fmt.Println(mySlice)

	fmt.Printf("%T\n", mySlice2)
	fmt.Printf("%v\n", mySlice2)
	fmt.Println(mySlice2)

	fmt.Printf("%T\n", mySlice3)
	fmt.Printf("%v\n", mySlice3)
	fmt.Println(mySlice3)

}
