package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

func init () {
	runtime.GOMAXPROCS(runtime.NumCPU())

}

func main () {
	cars := []string {"City", "Everest", "Civic", "Fortuner", "Accent"}

	wg.Add(2)

	foo(cars...)
	bar(cars...)

	wg.Wait()

}

func foo (car ...string) {
	for _, variant := range car {
		fmt.Println("Foo:\t" + variant)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()
}

func bar (car ...string) {
	for _, variant := range car {
		fmt.Println("Bar:\t" + variant)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}

	wg.Done()
}