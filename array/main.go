package main

import "fmt"

func main() {
	var x [58]string
	y := make([]string, 100)
	for i := 65; i < 122; i++ {
		x[i - 65] = string(i)	
	}

	for i := 65; i < 122; i++ {
		y[i - 65] = string(i)	
	}

	fmt.Println(x)
	fmt.Println(y)
}