package main

import (
	"fmt"
	"time"
)

func main () {
	c := make (chan int)

	go func () {
		for i := 0; i < 10; i++ {
			c <- 1
		}
	} ()

	go func () {
		for {
			fmt.Println(<-c)
		}
	} ()

	time.Sleep(time.Second)

}