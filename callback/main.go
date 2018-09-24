package main

import "fmt"

func main () {
	
	rnum := getNumber([]int {1,2,3,4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(rnum)
}

func getNumber(numbers []int, callback func(int) bool) []int {
	xs := []int{}

	for _,  num := range numbers {
		if callback(num) {
			xs = append(xs,num)
		}
	}

	return xs
}