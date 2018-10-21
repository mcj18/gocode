package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main () {
	wg.Add (2)
	go incrementor("Foo:\t")
	go incrementor("Bar:\t")
	wg.Wait ()
	fmt.Println ("Final Counter:\t", counter)

}

func incrementor(n string) {
	for i := 0; i < 20; i++ {
		time.Sleep (time.Duration (rand.Intn (20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(n, i, "Counter:\t", counter)
		mutex.Unlock()
	}

	wg.Done()

}