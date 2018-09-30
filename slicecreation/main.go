package main

import "fmt"

func main() {
	/*Initialized with nil*/
	var engineer []string
	var engineers [][]string

	/*Initialize with value*/
	eng := []string{"Michael Christian Jaramilla"}
	engs := [][]string{}

	/*using make*/
	en := make([]string, 100)
	ens := make([][]string, 100)

	engineer = append(engineer,"Michael")

	fmt.Println(engineer)
	fmt.Println(engineers)
	fmt.Println(engineers == nil)
	eng = append(eng, "Cherry")
	engs = append(engs, eng)
	fmt.Println(eng)
	fmt.Println(engs)
	fmt.Println(eng == nil)
	fmt.Println(en)
	fmt.Println(ens)
}