package main

import (
	//"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("firing all three concurrent requests")
	// MakeRequest(1)
	// MakeRequest(2)
	// MakeRequest(3)
	var wg sync.WaitGroup

	for i :=0; i<3; i++ {
		wg.Add(1)
		go func (id int)  {
			fmt.Printf("req (%d) starts a request\n", id)
			time.Sleep(3 * time.Second) //simulate a slow network traffic
			fmt.Printf("req (%d) done\n", id)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}

func MakeRequest(id int) {
	fmt.Printf("req (%d) makes the request\n", id)
	time.Sleep(3 * time.Second) //simulate a slow network that takes 3 secs
	fmt.Printf("req (%d) done\n", id)
}