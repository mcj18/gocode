package main

import "fmt"

var name string

func main() {

	name = "Michael"
	fmt.Println("Hello World!")
	fmt.Println("My name is " + name)

	if name == "Michael" {
		fmt.Println("True")
	} else {
		fmt.Println("False")

	}

	b:= true
	if name := "Alexander"; b {

		fmt.Println(name)
	} else {

		fmt.Println(name)
	}

	randomName := "Alexander" 

	if name := randomName; name == "Michael" {

		fmt.Println("Hello " + name + "!")
	} else if name == "Alexander" {

		fmt.Println("Whats up " + name + "?")
	} else {
		fmt.Println("It was nice meeting you. What's your name?")
		fmt.Println("I am " + name)

	}
}