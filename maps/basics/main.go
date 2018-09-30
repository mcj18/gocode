package main

import "fmt"

func main() {
	m := make(map[string] int)

	m["k1"] = 3
	m["k2"] = 15
	delete(m, "k2")
	fmt.Println(m)

	v,ok :=m["k1"]
	fmt.Println("Ok?", ok, v) 

}