package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int //shared state
	mu		sync.Mutex
)
/*
some operations might be discarded when trying to 
update a shared state
*/

func increment() {
	mu.Lock()
	defer mu.Unlock()
	value := counter
	//this takes a delay of 1 sec
	time.Sleep(time.Microsecond) //simulate a db query
	fmt.Println("v", value)
	counter = value + 1
}

func main() {
	expected := 1000 //this depicts 1000 counters
	var wg sync.WaitGroup
	//
	for i := 0; i<expected; i++{
		wg.Add(1)
		go func()  {
			increment()
			defer wg.Done()	
		}()
	}
	wg.Wait()
	fmt.Println("counter", counter)
	fmt.Println("expected", expected)
}