package main

import (
	"fmt"
)

func main() {
	name := "Michael"
	
	switch name {
		case "Joan": {
			fmt.Println("Hello" + name)
		}

		case "Cherry": {
			fmt.Println("How are you doin" + name)
		}

		case "Michael": {
			fmt.Println("It was nice meeting you " + name)
		}

		default:
			fmt.Println("unknown!")

	}
}