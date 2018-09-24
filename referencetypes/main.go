package main

import "fmt"

var number int

func main() {
	m := make([]string, 5, 10)
	s := make([] int, 1, 10)

	fmt.Println(m)
	fmt.Println(s)
	changeMe(m)
	changeNum(s)
	fmt.Println(m)
	fmt.Println(s)
}

func changeMe(n []string) {
	n[0] = "Michael"

	fmt.Println(n)
}

func changeNum(n []int) {
	n[0] = 5

	fmt.Println(n)
}

