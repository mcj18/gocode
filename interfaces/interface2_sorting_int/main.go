package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{5, 4, 3, 2, 1}

	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
}
