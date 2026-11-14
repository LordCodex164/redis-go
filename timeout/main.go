package main

import (
	//"context"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("firing all three concurrent requests")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	//we can use context to usually ``
	// MakeRequest(1)
	// MakeRequest(2)
	// MakeRequest(3)
	var wg sync.WaitGroup

	// for i :=0; i<3; i++ {
		// ``	wg.Add(1)
		// 	go func (id int)  {
		// 		// fm``t.Printf("req (%d) starts a request\n", id)
	// 		// time.Sleep(3 * time.Second) //simulate a slow network traffic
	// 		// fmt.Printf("req (%d) done\n", id)
	// 		MakeRequest(ctx, id)
	// 		defer wg.Done()
	// 	}(i)
	// }
	//make 5 concurrent requests

	for i :=0; i<5; i++ {
		wg.Add(1)
		go func (id int)  {
			MakeRequest(ctx, id)
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}

func MakeRequest(ctx context.Context, id int) {
	// fmt.Printf("req (%d) makes the request\n", id)
	// time.Sleep(3 * time.Second) //simulate a slow network that takes 3 secs
	// fmt.Printf("req (%d) done\n", id)
	
	

}