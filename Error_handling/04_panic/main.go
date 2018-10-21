package main

import (
	"os"
	"log"
)

func main () {
	_, err := os.Open ("no-file.txt")

	if err != nil {
		//Fatalln is equiv to Println() followed by a call to os.Exit()
		log.Fatalln(err)
	}
}