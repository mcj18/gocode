package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
