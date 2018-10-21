package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main () {
	name := []string {"michael", "cherry", "alexander", "edrei"}

	wg.Add(2)

	go foo (name...)
	go bar (name...)

	wg.Wait()

}

func foo (n ...string) {
	for _, v := range n {
		fmt.Println("Foo:\t" + v)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()

}

func bar (n ...string) {
	for _, v := range n {
		fmt.Println("Bar:\t" + v)
		time.Sleep(time.Duration(2 * time.Millisecond))
	}

	wg.Done()

}