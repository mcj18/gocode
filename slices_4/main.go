package main

import (
	"fmt"
)

func main() {

	mySlice1 := []int{1, 2, 3, 4, 5}
	mySlice2 := []int{6, 7, 8, 9, 10}
	var mySlice3 []int = nil
	mySlice4 := []int{1, 2, 3, 4, 5}
	mySlice5 := make([]int, 5)
	mySlices := make([][]int, 5, 10)
	mySlicesContainer := make([][]int, 5, 10)
	//var mySlices [][]int
	//mySlices := [][]int {{1,3,4,5,6}}
	fmt.Printf("%v\n", mySlice1)
	fmt.Printf("%v\n", mySlice2)
	fmt.Printf("%v\n", mySlice3)
	fmt.Printf("%v\n", mySlice4)
	fmt.Printf("%v\n", mySlice5)
	fmt.Printf("%v\n", mySlices)

	mySlices[0] = mySlice1
	mySlices[1] = mySlice2
	mySlices[2] = mySlice3
	mySlices[3] = mySlice4
	mySlices[4] = mySlice5

	fmt.Printf("%v\n", mySlices)

	//******Put data on a separate container[][]
	fmt.Println("Put data on a separate container[][]")
	fmt.Printf("%v\n", mySlicesContainer)
	for i, v := range mySlices {
		mySlicesContainer[i] = v
	}

	fmt.Printf("%v\n", mySlicesContainer)
}
